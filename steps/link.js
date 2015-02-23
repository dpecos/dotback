var fs = require("fs");

module.exports = function(ctx) {
   return function(bundle, file) {
      var dest = ctx.HOME;
      var source = ctx.DOTFILES + bundle;

      if (typeof(file) === 'object') {
         if (file.dest) {
            dest = dest + file.dest + "/";
         }
         file = file.files;
      }

      var processFile = function(fileSource, fileDest, remove) {
         if (remove) {
            if (fs.existsSync(fileDest)) {
               ctx.executeAction("rm " + fileDest, function() {
                  fs.unlinkSync(fileDest);
               });
            }
         } else {
            ctx.executeAction("link " + fileSource + " -> " + fileDest, function() {
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
   };
};
