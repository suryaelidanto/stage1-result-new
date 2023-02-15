// What is async await, the point is that async-await is simply other way to call promise but it looks like in synchronous way.

let condition = true;

let promise = new Promise(function (resolve, reject) {
    if (condition) {
        resolve("Promise is resolved");
    } else {
        reject("Promise is rejected");
    }
})

// if we previously call the promise like this :

// promise.then(function (value) {
//     console.log(value);
// }).catch(function (reason) {
//     console.log(reason)
// })

// now, with async-await, we can call it like this :

async function asyncFunction() {
    const response = await promise;
    console.log(response)
}

asyncFunction() // we call the function here

// it is looking more clean and easy to read right? 😁, using async-await considered as best practice.



