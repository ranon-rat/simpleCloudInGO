console.log("hello world");
interface Api {
  files: { id: number; filename: string }[];
  size: number;
}
fetch("/api/" + window.location.pathname.slice(1))
  .then((r) => r.json())
  .then((d: Api) => {
    d.files.forEach((i) => {
 
      document.getElementById("files").innerHTML+=`
      <div class="file">
        <a href="/getFile/${i.id}/${i.filename}">
         <div class="id">${i.id}</div> <div class="filename">${i.filename}</div>
          </div>
         </a>
      </div>`
    });
    

  });
