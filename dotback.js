#!/usr/bin/env node 
var fs = require("fs");
var path = require("path");
var child_process = require("child_process");
var exec = child_process.execSync;
var _ = require("lodash");
var os = require("os");
var log = require("winston");
log.cli()
log.level = 'debug';

var HOME = process.env.HOME + "/";
var DOTFILES = HOME + ".dotfiles/";

function parseRecipe(bundle, recipe) {
   var actions = [];
   var stepsCache = {};

   for (var action in recipe) {
      if (action == "skip") {
         return [];
      }
      if (!stepsCache[action]) {
         try {
         var step = require("./steps/" + action + ".js"); 
         stepsCache[action] = step({
            executeAction: executeAction,
            HOME: HOME,
            DOTFILES: DOTFILES,
            log: log
         });
         } catch (err) {
            log.error("Step '" + action + "' definition not found");
         }
      }

      var handler = stepsCache[action];

      var params = recipe[action];
      if (!Array.isArray(params)) {
         params = [ params ]
      }
      params.reverse();
      params.push(bundle);
      params.reverse();

      actions.push(handler.apply(this, params));
   }

   return actions;
}

function loadConfig(hostname) {
   var jsonConfig = null;
   if (hostname) {
      jsonConfig = DOTFILES + "config." + hostname + ".json";
   } else {
      jsonConfig = DOTFILES + "config.json";
   }

   if (fs.existsSync(jsonConfig)) {
      jsonConfig = require(jsonConfig);
      var config = {};

      for (var bundle in jsonConfig) {
         var recipe = jsonConfig[bundle];

         if (Array.isArray(recipe)) {
            config[bundle] = _(recipe).map(function(step) {
               return parseRecipe(bundle, step);
            }).flatten().value();
         } else {
            config[bundle] = parseRecipe(bundle, recipe); 
         }
      }
      return config;
   } else {
      return null;
   }
}

var debugEnabled = false;

function executeAction(message, action) {
   log.debug(message);
   if (!debugEnabled) {
      if (action) {
         action();
      }
   }
}

function debug(enable) {
   debugEnabled = enable;
}

function clean() {
   perform(true);
}

function setup() {
   perform(false);
}

function perform(remove) {
   for (key in config) {
      if (config.hasOwnProperty(key)) {
         if (Array.isArray(config[key])) {
            config[key].forEach(function(action) {
               action(remove);
            });
         } else {
            config[key](remove);
         }
      }
   }
}

function init(repository) {
   if (repository != null) {
      exec("git clone " + repository + " " + DOTFILES);
   } else {
      if (!fs.existsSync(DOTFILES)) {
         fs.mkdir(DOTFILES);
      }
      if (!fs.existsSync(DOTFILES + "/.git")) {
         exec("cd " + DOTFILES + " && git init");
      }
   }
}

var argv = require("optimist").
   usage("Usage: $0 --action [action] [--debug]").
   demand("action").string("action").describe("action", "Possible values: init, install").
   argv;


if (argv.action == "init") {

   if (argv._.length != 1) {
      log.error("init requires a repository URL as single parameter");
   } else {
      init(process.argv[3]);
   }

} else if (argv.action == "install") {

   if (argv.debug) {
     log.info("Debug mode enabled");
     debug(true);
   }

   var config = loadConfig();
   var hostConfig = loadConfig(os.hostname());

   config = _.extend(config, hostConfig);

   log.info("Cleaning files...");
   clean();
   log.info("Setting up environment...");
   setup();

} else {
   log.error("Unknown action " + argv.action);
}
