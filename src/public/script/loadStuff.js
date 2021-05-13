console.log("hello world");
fetch("/api/" + window.location.pathname.slice(1))
    .then(function (r) { return r.json(); })
    .then(function (d) {
    d.files.forEach(function (i) {
        document.getElementById("files").innerHTML += "\n      <p class=\"file\">\n        <a href=\"/getFile/" + i.id + "/" + i.filename + "\">\n          <div  style=\"background-image:url('/getFile/" + i.id + "/" + i.filename + "')\" class=\"image-file\">\n          </div>\n         </a>\n      </p>";
    });
});
