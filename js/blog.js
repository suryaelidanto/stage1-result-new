// let namaSiswa1 = "Bagus"
// let namaSiswa2 = "Rizal"
// let namaSiswa3 = "Khalid"

// console.log(namaSiswa1);
// console.log(namaSiswa2);
// console.log(namaSiswa3);

// // Array
// let namaSiswa = ["Bagus", "Rizal", "Khalid", 2022, true]
// console.log(namaSiswa);
// console.log(namaSiswa[1]);
// console.log(namaSiswa[2]);

// Object
// let nama = "Bagus"
// let alamat = "Tangerang"
// let umur = 20

// console.log(nama);
// console.log(alamat);
// console.log(umur);

// let dataPersonal1 = {
//     nama: "Bagus",
//     alamat: "Tangerang",
//     umur: 20,
// }

// let dataPersonal2 = {
//     nama: "Rizal",
//     alamat: "Jakarta",
//     umur: 20,
// }
// let dataPersonal3 = {
//     nama: "Rizal",
//     alamat: "Jakarta",
//     umur: 20,
// }

// let dataPersonal4 = {
//     nama: "Rizal",
//     alamat: "Jakarta",
//     umur: 20,
// }

// console.log(dataPersonal1);
// console.log(dataPersonal1.alamat);

// Array Object
// let dataPersonal = [
//     {
//         nama: "Aziz",
//         alamat: "Jakarta"
//     },
//     {
//         nama: "Rizal",
//         alamat: "Jakarta",
//     },
//     {
//         nama: "Khalid",
//         alamat: "Jakarta",
//     }
// ]

// console.log(dataPersonal);

// let data = []

// function addData(){

//     let person = {
//         name: "Samsul",
//         address: "Tangerang"
//     }
    
//     data.push(person)
//     console.log(data);
// }

// let data = []

// function addData(){

//     let blog = {
//         title: document.getElementById("input-blog-title").value,
//         content: document.getElementById("input-blog-content").value
//     }

//     data.push(blog)
//     console.log(data);
// }


let dataBlog = []
function addBlog(event) {
    event.preventDefault()

    let title = document.getElementById("input-blog-title").value
    let content = document.getElementById("input-blog-content").value
    let image = document.getElementById("input-blog-image").files

    // untuk membuat url gambar, agar tampil
    image = URL.createObjectURL(image[0])
    console.log(image);

    let blog = {
        title,
        content,
        image,
        postAt: "15 September 2022",
        author: "Surya Elidanto"
    }

    dataBlog.push(blog)
    console.log(dataBlog);

    renderBlog()
}


function renderBlog(){

    document.getElementById("contents").innerHTML = ''

    for (let index = 0; index < dataBlog.length;  index++){  
        console.log(dataBlog[index]);

        document.getElementById("contents").innerHTML += `
        
            <div class="blog-list-item">
                <div class="blog-image">
                <img src="${dataBlog[index].image}" alt="" />
                </div>
                <div class="blog-content">
                <div class="btn-group">
                    <button class="btn-edit">Edit Post</button>
                    <button class="btn-post">Delete Post</button>
                </div>
                <h1>
                    <a href="blog-detail.html" target="_blank"
                    >${dataBlog[index].title}</a
                    >
                </h1>
                <div class="detail-blog-content">
                    ${dataBlog[index].postAt} | ${dataBlog[index].author}
                </div>
                <p>
                    ${dataBlog[index].content}
                </p>
                </div>
            </div>
        `
    }
}



// document.getElementById("card").innerHTML = ''

// for (let index = 0; index < dataBlog.length;  index++){  
//     console.log(dataBlog[index]);

//     document.getElementById("card").innerHTML += `
    
//     <div style="width: 400px; margin: 10px;">
//         <img src="${dataBlog[index].image}" alt="" width="100%">
//         <h2>${dataBlog[index].title}</h2>
//     </div>
//     `
// }
// }