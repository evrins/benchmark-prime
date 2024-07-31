const isPrime = (n) => {
  if (n < 2) {
    return false;
  }

  if (n === 2) {
    return true;
  }

  const max = Math.ceil(Math.sqrt(n));
  for (let i = 2; i <= max; i++) {
    if (n % i === 0) {
      return false;
    }
  }

  return true;
};

const main = (n) => {
  let count = 0;
  const startTime = Date.now();
  for (let i = 0; i < n; i++) {
    if (isPrime(i)) {
      count += 1;
    }
  }

  const endTime = Date.now();
  console.log(`node, ${n}, ${count}, ${endTime - startTime}`)
};

const args = process.argv;
if (args.length !== 3) {
  console.log(`Usage: ${args[1]} <n>`);
  process.exit(1);
}

main(args[2]);
