function hello() {
    console.log('Hello, world!');
}

function goodbye() {
    console.log('Goodbye, world!');
}

function print(callback) {
    callback();
}

print(hello);
print(goodbye);