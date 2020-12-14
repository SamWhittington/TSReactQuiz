let raceNumber = Math.floor(Math.random() * 1000);

const registeredEarly = false;
const age = 17;

if (age >= 18 && registeredEarly) {
  raceNumber += 1000
}

if (age > 18 && registeredEarly) {
  console.log("Racer number: " + raceNumber + " your race will begin at 9:30")
} else if (age > 18 && !registeredEarly){
   console.log("Racer number: " + raceNumber + " your race will begin at 11:00")
} else if (age < 18) { 
  console.log("Racer number: " + raceNumber + " your race will begin at 12:30")
}
