var child_process = require("child_process");
var exec = child_process.execSync;

module.exports = function(ctx) {
   return function(bundle, cmd) {
      return function(remove) {
         if (!remove) {
            var command = cmd;
            if (typeof(cmd) === 'object') {
               command = cmd.cmd;
               if (cmd.cwd) {
                  command = "cd " + cmd.cwd + " && " + command;
               }
            }
            if (command) {
               ctx.executeAction(command, function() {
                  try {
                     exec(command);
                  } catch (err) {
                     ctx.log.error("EXEC: error executing '" + command + "': " + err);
                  }
               });
            }
         }
      };
   }
};
