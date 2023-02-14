// Polymorphism
class Vehicle {
    drive() {
        return "The vehicle is driving.";
    }
}

class Car extends Vehicle {
    drive() {
        return "The car is driving.";
    }
}

class ElectricCar extends Car {
    drive() {
        return "The electric car is driving silently.";
    }
}

let myVehicle = new Vehicle();
let myCar = new Car();
let myElectricCar = new ElectricCar();

console.log(myVehicle.drive()); // Output: The vehicle is driving.
console.log(myCar.drive()); // Output: The car is driving.
console.log(myElectricCar.drive()); // Output: The electric car is driving silently.
