var exec = require("exec-sync");

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
            ctx.executeAction(command, function() {
               try {
                  exec(command);
               } catch (err) {
                  console.log(err);
               }
            });
         }
      };
   }
};
