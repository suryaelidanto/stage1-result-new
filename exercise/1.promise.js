
let condition = true;

let promise = new Promise(function (resolve, reject) {
    if (condition) {
        resolve("Promise is resolved");
    } else {
        reject("Promise is rejected");
    }
})

// console.log(promise)
// what? why it is appear Promise? because we need to wait it, that's the point of the promise, we should access it like this :
promise.then(function (value) {
    console.log(value);
}).catch(function (reason) {
    console.log(reason)
})


