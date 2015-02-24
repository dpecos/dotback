var chai = require("chai");
var sinon = require("sinon");
var chaiAsPromised = require("chai-as-promised");
var sinon_chai = require("sinon-chai");
var fs = require("fs");

chai.use(chaiAsPromised);
chai.use(sinon_chai);

var should = chai.should();

describe("git step", function() {
   var git = null;
   var action = null;
   var test_dir = "/tmp/.test-git"

   before(function() {
      var step = require("../steps/git.js");
      git = step({
         HOME: "/tmp/",
         executeAction: function(command, fn) {
            action = {
               command: command
            };
         }
      });
   });

   beforeEach(function() {
      action = null;

      fs.mkdirSync(test_dir);
   });

   afterEach(function() {
      fs.rmdirSync(test_dir);
   });

   it("should clone a repo", function() {
      var step = git("test-git", "repository_url");
      step();

      action.command.should.equal("git clone repository_url /tmp/.test-git");
   });

   it("should clean its actions if file exist", function() {
      var step = git("test-git", "repository_url");
      step(true);

      action.command.should.equal("rm -r /tmp/.test-git");
   });

   it("should not clean its actions if file doesn't exist", function() {
      var step = git("dont-exists-directory", "repository_url");
      step(true);

      (action.command === null).should.be.true;
   });

});
