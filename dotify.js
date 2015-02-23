#!/usr/bin/env node 
var fs = require("fs");
var path = require("path");
var exec = require('exec-sync');
var _ = require("lodash");

var HOME = process.env.HOME + "/";
var DOTFILES = HOME + ".dotfiles/";

function parseRecipe(bundle, recipe) {
   var actions = [];

   for (var action in recipe) {
      if (action == "skip") {
         return [];
      }
      var handler = global[action];

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

function loadConfig() {
   var jsonConfig = require(DOTFILES + "/config.json");
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
}

global.link = function(bundle, file) {
   var dest = HOME;
   var source = DOTFILES + bundle;

   if (typeof(file) === 'object') {
      if (file.dest) {
         dest = dest + file.dest + "/";
      }
      file = file.files;
   }

   var processFile = function(fileSource, fileDest, remove) {
      if (remove) {
         if (fs.existsSync(fileDest)) {
            executeAction("rm " + fileDest, function() {
               fs.unlinkSync(fileDest);
            });
         }
      } else {
         executeAction("link " + fileSource + " -> " + fileDest, function() {
            fs.symlinkSync(fileSource, fileDest); 
         });
      }
   };

   if (file === null || file !== "*") {
      return function(remove) {
         var fileDest = dest + "." + bundle;
         var fileSource = source;
         if (file) {
            fileSource = fileSource + "/" + file;
            fileDest = dest + "." + file;
         }

         processFile(fileSource, fileDest, remove);
      }
   } else if (file === "*") {
      return function(remove) {
         var files = fs.readdirSync(source);
         files.forEach(function(file) {
            var fileSource = source + "/" + file;
            var fileDest = dest + "." + file;

            processFile(fileSource, fileDest, remove);
         });
      }
   }
}

global.exec = function(bundle, cmd) {
   return function(remove) {
      if (!remove) {
         var command = cmd;
         if (typeof(cmd) === 'object') {
            command = cmd.cmd;
            if (cmd.cwd) {
               command = "cd " + cmd.cwd + " && " + command;
            }
         }
         executeAction(command, function() {
            try {
               exec(command);
            } catch (err) {
               console.log(err);
            }
         });
      }
   };
}

global.git = function(orig, repo) {
   var dest = HOME + "." + orig;
   return function(remove) {
      var command = null;
      if (remove) {
         if (fs.existsSync(dest)) {
            command = "rm -r " + dest;

         }
      } else {
         command = "git clone " + repo + " " + dest;
      }
      executeAction(command, function() {
         try {
            exec(command);
         } catch (err) {
            console.log(err);
         }
      });
   };
}

var debugEnabled = false;

function executeAction(message, action) {
   if (debugEnabled) {
      console.log(message);
   } else {
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

if (process.argv[2] == "init") {
   init(process.argv[3]);

} else if (process.argv[2] == "install") {
   var config = loadConfig();

   if (process.argv[3] == "--debug") {
     debug(true);
   }

   clean();
   setup();
} else {
   console.log("ERROR: Unknown option");
}
