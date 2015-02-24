var chai = require("chai");
var sinon = require("sinon");
var chaiAsPromised = require("chai-as-promised");
var sinon_chai = require("sinon-chai");

chai.use(chaiAsPromised);
chai.use(sinon_chai);

var should = chai.should();

describe("git step", function() {
   var git = null;
   var action = null;

   before(function() {
      var step = require("../steps/git.js");
      git = step({
         HOME: process.env.HOME + "/",
         executeAction: function(command, fn) {
            action = {
               command: command
            };
         }
      });
   });

   beforeEach(function() {
      action = null;
   });

   it("should clone a repo", function() {
      var step = git("orig", "repository_url");
      step();

      action.command.should.equal("git clone repository_url " + process.env.HOME + "/.orig");
   });

   it("should clean its actions if file exist", function() {
      var step = git("config", "repository_url");
      step(true);

      action.command.should.equal("rm -r " + process.env.HOME + "/.config");
   });

   it("should not clean its actions if file doesn't exist", function() {
      var step = git("dont-exists-directory", "repository_url");
      step(true);

      (action.command === null).should.be.true;
   });

});
