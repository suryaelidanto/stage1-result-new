const promise = new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest();
    xhr.open('GET', 'https://api.npoint.io/3b94410b2076e1697d68', true);
    // console.log(xhr)
    xhr.onload = () => {
        if (xhr.status === 200) {
            // we parsing so it is easier to read in console
            // response vs reponseText, the difference is responseText older version, response is newer version, but for our course right now it is just the same.
            resolve(JSON.parse(xhr.response));
        } else {
            reject('Error loading data');
        }
    };
    xhr.onerror = () => {
        reject('Network error');
    };
    xhr.send();
})

async function getAllTestimonials() {
    const response = await promise;
    console.log(response)

    let testimonialHTML = '';
    const testimonial = response.map(function (item) {
        testimonialHTML += `<div class="testimonial">
            <img src="${item.image}" class="profile-testimonial" />
            <p class="quote">"${item.quote}"</p>
            <p class="author">- ${item.author}</p>
            <p class="author">${item.rating} <i class="fa-solid fa-star"></i></p>
        </div>`
    })

    document.getElementById('testimonials').innerHTML = testimonialHTML;
}

getAllTestimonials()

async function getFilteredTestimonials(rating) {
    const response = await promise;

    const testimonialFiltered = response.filter(function (item) {
        return item.rating === rating
    })

    console.log(testimonialFiltered)

    let testimonialHTML = '';

    if (testimonialFiltered.length === 0) {
        testimonialHTML = '<h1> Data not found! </h1>';
    } else {
        const testimonial = testimonialFiltered.map(function (item) {
            testimonialHTML += `<div class="testimonial">
                <img src="${item.image}" class="profile-testimonial" />
                <p class="quote">"${item.quote}"</p>
                <p class="author">- ${item.author}</p>
                <p class="author">${item.rating} <i class="fa-solid fa-star"></i></p>
            </div>`
        })
    }

    document.getElementById('testimonials').innerHTML = testimonialHTML;
}