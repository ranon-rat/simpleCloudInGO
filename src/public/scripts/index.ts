interface IApi {
  files: {
    id: number,
    filename: string;
  }[];
  size: number;
};

(async () => {
  try {
    const response: Response = await fetch("/api/" + window.location.pathname.slice(1));
    const res: IApi = await response.json();

    for (const i of res.files) {
      const extention: string[] = i.filename.split('.');
      let extTarget: string = 'default';

      for (const j of extention) {
        try {
          const resp: Response = await fetch(`/public/images/ext/${j}.svg`);
          if (resp.ok) extTarget = j;
        } catch (err: unknown) {};
      };

      document.getElementById("fcontainer")!.innerHTML += //html
      `<a href="/getFile/${i.id}/${i.filename}">
        <div class="fileSeparator">
          <div class="fileIImage">
            <img class="fileImage" src="/public/images/ext/${extTarget}.svg" />
          </div>
          <div class="fileName">
            <p>${i.filename}</p>
          </div>
          <div class="fileButton">
            <img class="fileBImage" src="/public/images/download.svg" />
          </div>
        </div>
      </a>`;
    };
  } catch (err: unknown) {
    console.log(err);
  };
})();