console.log("hello world");
fetch("/api/" + window.location.pathname.slice(1))
    .then(function (r) { return r.json(); })
    .then(function (d) {
    d.files.forEach(function (i) {
        document.getElementById("files").innerHTML += "\n      <p class=\"file\">\n        <a href=\"/getFile/" + i.id + "/" + i.filename + "\">\n         <span class=\"id\">" + i.id + "</span> <span class=\"filename\">" + i.filename + "</span>\n          </div>\n         </a>\n      </p>";
    });
});
