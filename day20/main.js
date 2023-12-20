import fs from "fs";

const inputPath = "./day20/example-input";

const MODULE_TYPES = {
  BROADCASTER: "b",
  FLIP_FLOP: "%",
  CONJUNCTION: "&",
};

const PULSES = {
  HIGH: 1,
  LOW: 0,
};

export function partOne(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  const modulesByKey = input
    .trim()
    .split("\n")
    .reduce((accum, line) => {
      let [module, destModules] = line.split(" -> ");
      console.log({ module, destModules });
      if (module === "broadcaster") {
        accum[module] = {
          type: MODULE_TYPES.BROADCASTER,
          outputs: destModules.split(", "),
          state: null,
        };
      } else {
        accum[module.slice(1)] = {
          type: module.slice(0, 1),
          outputs: destModules.split(", "),
          state: null,
          memory: null,
        };
      }
      return accum;
    }, {});

  console.log(modulesByKey);

  // Start the queue with a button press
  const queue = [{ type: PULSES.LOW, destinationKey: "broadcaster", fromKey: null }];
  // Start with a list of all modules being set to 'off' (0)
  const moduleStates = {};
  for (const moduleKey in modulesByKey) {
    moduleStates[moduleKey] = 0;
  }

  let pulses = { [PULSES.HIGH]: 0, [PULSES.LOW]: 0 };

  let presses = 0;
  while (presses < 1000) {
    // This is cyclic, so we don't need to do this 1000 times. Instead, find when it repeats and look at the modulo again
    while (queue.length) {
      console.log({ queue });
      let next = queue.shift();
      sendPulse(next.type, next.destinationKey, next.fromKey);
    }

    if (Object.values(moduleStates).every((val) => val === 1)) {
      // We're back at the start!
      debugger;
      break;
    }
    presses++;
  }

  console.log({ day: 20, part: 1, value: presses });

  // Now start the cycle by pressing the button, which sends a LOW PULSE to broadcast

  function sendPulse(pulseType, destinationKey, inputKey) {
    pulses[pulseType]++;
    // Bunch of if statements to handle each one.
    const module = modulesByKey[destinationKey];

    if (module.type === MODULE_TYPES.BROADCASTER) {
      // Send to all children by adding them to the queue in order
      addToQueue(pulseType, module.outputs, destinationKey);
    } else if (module.type === MODULE_TYPES.FLIP_FLOP) {
      // Flip-flop modules we only really care about low pulse being recieved
      if (pulseType === PULSES.LOW) {
        // If currently on, turn off and send low pulse
        if (module.state) {
          module.state = 0;
          addToQueue(PULSES.LOW, module.outputs, destinationKey);
        } else {
          // If off, turn on and send high pulse
          module.state = 1;
          addToQueue(PULSES.HIGH, module.outputs, destinationKey);
        }
      }
    } else if (module.type === MODULE_TYPES.CONJUNCTION) {
      // Conjunction modules (prefix &) remember the type of the most recent pulse received from each of their connected input modules;
      // they initially default to remembering a low pulse for each input. When a pulse is received, the conjunction module first updates
      // its memory for that input. Then, if it remembers high pulses for all inputs, it sends a low pulse; otherwise, it sends a high
      // pulse.

      // If null, then we use defaults of low for each input.
      // First time we see this node, determine inputs and set up the memory
      if (module.memory === null) {
        module.memory = {};
        // First build input module if doesnt exist
        for (const key in modulesByKey) {
          if (key === destinationKey) {
            continue;
          }
          const element = modulesByKey[key];
          if (element.outputs.includes(destinationKey)) {
            module.memory[key] = PULSES.LOW;
          }
        }
      }
      if (inputKey) {
        module.memory[inputKey] = pulseType;
      }

      // If ALL are high, send low, else send high
      addToQueue(
        Object.values(module.memory).some((el) => el === PULSES.LOW) ? PULSES.HIGH : PULSES.LOW,
        module.outputs,
        destinationKey
      );
    } else {
      throw Error("Module type not implemented?");
    }
  }

  function addToQueue(type, destinationKeys, fromKey) {
    destinationKeys.forEach((destinationKey) => {
      queue.push({ type, destinationKey, fromKey });
    });
  }
}

export function partTwo(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  var inputArray = input
    .trim()
    .split("\n")
    .map((line) => line.split(""));

  console.log({ day: 20, part: 2, value: "todo" });
}
