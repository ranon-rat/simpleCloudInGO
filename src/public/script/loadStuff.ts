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
      <p class="file">
        <a href="/getFile/${i.id}/${i.filename}">
         <span class="id">${i.id}</span> <span class="filename">${i.filename}</span>
          </div>
         </a>
      </p>`
    });
    

  });
