window.addEventListener('load', function () {
  setTimeout(function () {
    var webIamRoot = document.querySelector('webiam-root');
    // Switch to plain old HTML page, if browser is too old to bootstrap Angular
    if (webIamRoot && webIamRoot.attributes && !webIamRoot.attributes.length) {
      if (document.documentElement.lang === 'fr') {
        location.assign('/unsupported-browser/fr.html');
      } else if (document.documentElement.lang === 'it') {
        location.assign('/unsupported-browser/it.html');
      } else {
        location.assign('/unsupported-browser/de.html');
      }
    }
  }, 3000);
});
