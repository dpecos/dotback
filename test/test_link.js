var chai = require("chai");
var sinon = require("sinon");
var chaiAsPromised = require("chai-as-promised");
var sinon_chai = require("sinon-chai");
var fs = require("fs");

chai.use(chaiAsPromised);
chai.use(sinon_chai);

var should = chai.should();

describe("link step", function() {
   var link = null;
   var action = null;

   var home = "/tmp/";
   var dotfiles = home + ".dotfiles/";
   var test_dir = dotfiles + "test/";

   before(function() {
      var step = require("../steps/link.js");
      link = step({
         HOME: home,
         DOTFILES: dotfiles,
         executeAction: function(command, fn) {
            action = {
               command: command
            };
         }
      });
      fs.mkdirSync(dotfiles);
   });

   after(function() {
      fs.rmdirSync(dotfiles);
   });

   beforeEach(function() {
      action = null;

      fs.mkdirSync(test_dir);
      fs.writeFileSync(test_dir + "file");
   });

   afterEach(function() {
      fs.unlinkSync(test_dir + "file");
      fs.rmdirSync(test_dir);
   });

   it("should link directory", function() {
      var step = link("test", null);
      step();

      action.command.should.equal("link " + dotfiles + "test -> " + home + ".test");
   });

   it("should link a single file in directory", function() {
      var step = link("test", "file");
      step();

      action.command.should.equal("link " + dotfiles + "test/file -> " + home + ".file");
   });

   it("should link all directory content", function() {
      var step = link("test", "*");
      step();

      action.command.should.equal("link " + dotfiles + "test/file -> " + home + ".file");
   });

   it("should link matching directory content", function() {
      var step = link("test", "f*");
      step();

      action.command.should.equal("link " + dotfiles + "test/file -> " + home + ".file");
   });
    
   it("should not link non-matching directory content", function() {
      var step = link("test", "x*");
      step();

      (action === null).should.be.true;
   });

   it("should remove previous links if it exist", function() {
      fs.writeFileSync(home + ".file");

      var step = link("test", "*");
      step(true);

      action.command.should.equal("rm " + home + ".file");

      fs.unlinkSync(home + ".file");
   });

   it("should not remove previous links if it doesn't exist", function() {
      var step = link("test", "not-existing-file");
      step(true);

      (action === null).should.be.true;
   });

});
