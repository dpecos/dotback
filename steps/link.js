var fs = require("fs");

module.exports = function(ctx) {
   return function(bundle, file) {
      var dest = ctx.HOME;
      var source = ctx.DOTFILES + bundle;

      if (file !== null && typeof(file) === 'object') {
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

      if (file === null || file.indexOf("*") === -1) {
         return function(remove) {
            var fileDest = dest + "." + bundle;
            var fileSource = source;
            if (file) {
               fileSource = fileSource + "/" + file;
               fileDest = dest + "." + file;
            }

            processFile(fileSource, fileDest, remove);
         }
      } else {
         return function(remove) {
            var files = fs.readdirSync(source);
            files.forEach(function(fileName) {
               var matches = true;

               if (file.indexOf("*") !== -1 && file !== "*") {
                  matches = fileName.match(file);
                  matches = ( matches != null && matches[0] !== "");
               }

               if (matches) {
                  var fileSource = source + "/" + fileName;
                  var fileDest = dest + "." + fileName;

                  processFile(fileSource, fileDest, remove);
               }
            });
         }
      }
   };
};
