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
          <div  style="background-image:url('/getFile/${i.id}/${i.filename}')" class="image-file">
          </div>
         </a>
      </p>`
    });
    

  });
