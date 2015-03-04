var fs = require("fs");
var child_process = require("child_process");
var exec = child_process.existsSync;

module.exports = function(ctx) {
   return function(orig, repo) {
      var dest = ctx.HOME + "." + orig;
      return function(remove) {
         var command = null;
         if (remove) {
            if (fs.existsSync(dest)) {
               command = "rm -r " + dest;
            }
         } else {
            command = "git clone " + repo + " " + dest;
         }
         ctx.executeAction(command, function() {
            try {
               exec(command);
            } catch (err) {
               console.log(err);
            }
         });
      };
   }
};
