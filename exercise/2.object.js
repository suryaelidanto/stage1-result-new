// Class
class Car {
    constructor(make, model) {
        this.make = make;
        this.model = model;
    }

    getInfo() {
        return `The car is a ${this.make} ${this.model}.`;
    }
}


// Object
let myCar = new Car("Toyota", "Camry");
console.log(myCar.getInfo()); // Output: The car is a Toyota Camry.