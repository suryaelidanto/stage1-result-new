// Abstraction:
class Car {
    constructor(make, model) {
        this._make = make;
        this._model = model;
    }

    get make() {
        return this._make;
    }

    get model() {
        return this._model;
    }
}

let myCar = new Car("Toyota", "Camry");
console.log(myCar.make); // Output: Toyota
console.log(myCar.model); // Output: Camry