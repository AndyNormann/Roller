let arg = process.argv[2].match(/(\d*)?d(\d*)\+?(\d*)?/)

let amount = arg[1]
let dice = arg[2]
let plus = +arg[3]

let rolls = []

let current = Math.floor(Math.random() * dice) + 1
rolls.push(current)
let ret_val = 0
ret_val += current

if(amount) {
  for(let i = 1; i < amount; i++) {
    current = Math.floor(Math.random() * dice) + 1
    rolls.push(current)
    ret_val += current
  }
}

if(plus) {
  rolls.push(plus)
  ret_val += plus
}

console.log("Rolls: ", rolls.join("+"));
console.log("Total: ", ret_val);
