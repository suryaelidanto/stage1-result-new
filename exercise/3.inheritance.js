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

// Inheritance
class ElectricCar extends Car {
    constructor(make, model, batteryCapacity) {
        super(make, model);
        this.batteryCapacity = batteryCapacity;
    }

    getInfo() {
        return `${super.getInfo()} It has a battery capacity of ${this.batteryCapacity} kWh.`;
    }
}


let myElectricCar = new ElectricCar("Tesla", "Model S", 100);
console.log(myElectricCar.getInfo()); // Output: The car is a Tesla Model S. It has a battery capacity of 100 kWh.
