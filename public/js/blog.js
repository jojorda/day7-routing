
let data = []

function addData(event) {

    event.preventDefault()

    let projecName = document.getElementById("input-blog-title").value;
    let starDate = document.getElementById("start-date").value;
    let endDate = document.getElementById("end-date").value;
    let decription = document.getElementById("input-blog-content").value;
    let image = document.getElementById("input-blog-image").files;

    let gambar = URL.createObjectURL(image[0])

    console.log("gambar", image[0])
    console.log("gambar dengan path", gambar)

    let blog = {
        projecName,
        starDate,
        endDate,
        decription,
        gambar,
        postAt: new Date(),
        author: "jordan Cool"
    }

    data.push(blog)
    console.log(data)
    renderBlog()
}

function renderBlog() {
    document.getElementById("contents").innerHTML = ``
    for (let index = 0; index < data.length; index++) {
        document.getElementById("contents").innerHTML += `<div class="blog-list-item">
    <div class="blog-image">
      <img src="${data[index].gambar}" alt="" />
    </div>
    <div class="blog-content">
      <div class="btn-group">
        <button class="btn-edit">Edit Post</button>
        <button class="btn-post">Post Blog</button>
      </div>
      <h1>
        <a href="blog-detail.html" target="_blank">${data[index].projecName}</a>
      </h1>
      <div class="detail-blog-content">
      ${data[index].starDate}-${data[index].endDate} | ${data[index].author}
      </div>
      <p>
      ${data[index].decription}
      </p>
    </div>
  </div>`
    }
}


