var chai = require("chai");
var sinon = require("sinon");
var chaiAsPromised = require("chai-as-promised");
var sinon_chai = require("sinon-chai");

chai.use(chaiAsPromised);
chai.use(sinon_chai);

var should = chai.should();

describe("exec step", function() {
   var exec = null;
   var action = null;

   before(function() {
      var step = require("../steps/exec.js");
      exec = step({
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

   it("should execute a simple echo command", function() {
      var cmd = 'echo "Hello World!"'; 
      var step = exec("tests", cmd);
      step();

      action.command.should.equal(cmd);
   });

   it("should execute a simple echo command in a set working directory", function() {
      var cmd = 'echo "Hello World!"'; 
      var path = "/path/to/tests";
      var step = exec("tests", {
         cmd: cmd,
         cwd: path
      });
      step();

      action.command.should.equal("cd " + path + " && " + cmd);
   });


});
