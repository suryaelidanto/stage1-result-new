// Encapsulation
class Car {
    constructor(make, model) {
        this._make = make;
        this._model = model;
    }

    get make() {
        return this._make;
    }

    set make(value) {
        if (value === "") {
            console.log("The make cannot be an empty string.");
            return;
        }
        this._make = value;
    }

    get model() {
        return this._model;
    }

    set model(value) {
        if (value === "") {
            console.log("The model cannot be an empty string.");
            return;
        }
        this._model = value;
    }

    drive() {
        console.log(`Driving a ${this._make} ${this._model}.`);
    }
}

let car = new Car("Toyota", "Camry");
console.log(car.make); // Output: Toyota
car.make = "Honda";
console.log(car.make); // Output: Honda

car.drive(); // Output: Driving a Honda Camry.


