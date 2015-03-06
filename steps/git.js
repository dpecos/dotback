var fs = require("fs");
var child_process = require("child_process");
var exec = child_process.execSync;

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

         if (command) {
            ctx.executeAction(command, function() {
               try {
                  exec(command);
               } catch (err) {
                  ctx.log.error("GIT: error executing '" + command + "': " + err);
               }
            });
         }
      };
   }
};
