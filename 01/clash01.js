const c = process.argv[2] // char to be inserted in the middle
const l = parseInt(process.argv[3]) // total length of the output

if (!c || c.length !== 1 || !l || isNaN(l)) {
  throw new Error('Expected types: arg1 char, arg2 int')
}

if (l % 2 === 0 || l < 2) {
  console.log("CAN'T")
} else {
  const pad = '*'.repeat(l / 2)
  console.log(`${pad}${c}${pad}`)
}
