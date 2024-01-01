import fs from "fs";

const inputPath = "./day25/example-input";

export function partOne(input = null) {
  var input = input || inputPath;
  input = fs.readFileSync(input, "utf8");
  const graph = buildGraph(input);

  /* Let's try to implement Karger's algorithm */
  // Pick a random edge and merge it's endpoints into a 'super-node'
  // Repeat until there are only two nodes left
  // Count the number of edges between the two nodes
  // Repeat the algorithm a number of times and until we get 3 (we know that's the minimum)

  // Build edges array - each is an array of two nodes src-dst (although it's not directional, I think this is fine?)
  const edges = [];
  for (let [key, value] of Object.entries(graph)) {
    for (let connection of value) {
      if (edges.some((edge) => edge.includes(connection) && edge.includes(key))) continue;
      edges.push([key, connection]);
    }
  }

  const V = Object.keys(graph).length;

  let cutEdges = 0;
  let cutEdgeKeys = [];
  let countInRoot = {};

  while (cutEdges !== 3) {
    [cutEdges, cutEdgeKeys, countInRoot] = kargerMinCut();
  }

  console.log({ day: 25, part: 1, value: Object.values(countInRoot).reduce((a, b) => a * b, 1) });

  function kargerMinCut() {
    let vertices = V;

    // Build the subSets - this is how we will merge the nodes
    let subSets = {};
    for (const key in graph) {
      // All start as their own parent - order 0 is the lowest, and larger means 'parent', 'grandparent', etc.
      subSets[key] = { parent: key, order: 0 };
    }

    while (vertices > 2) {
      // Pick a random edge:
      let edge = edges[Math.floor(Math.random() * edges.length)];
      let [src, dst] = edge;

      let srcRoot = find(src);
      let dstRoot = find(dst);

      if (srcRoot === dstRoot) continue;

      // Contract this edge - remove a vertex and handle the subset union
      handleMerge(src, dst);
      vertices--;
    }

    // Determine how many edges are cut on this run
    let cutEdges = 0;
    let cutEdgeKeys = [];
    for (let edge of edges) {
      let [src, dst] = edge;
      let srcRoot = find(src);
      let dstRoot = find(dst);
      // If they are from different roots, then this is a cut edge
      if (srcRoot !== dstRoot) {
        cutEdges++;
        cutEdgeKeys.push(edge);
      }
    }

    // We need the count of nodes in each root (to multiply them later), so let's build that and return it.
    let countInRoot = {};
    for (const key in subSets) {
      let rootFromNode = find(key);
      countInRoot[rootFromNode] = (countInRoot[rootFromNode] ?? 0) + 1;
    }

    return [cutEdges, cutEdgeKeys, countInRoot];

    /**
     * Finds the root of the subset that the given key is in
     */
    function find(key) {
      if (subSets[key].parent !== key) {
        subSets[key].parent = find(subSets[key].parent);
      }
      return subSets[key].parent;
    }

    /**
     * Handles the hierarchical merging of two subsets - one (of their roots) will become a parent of the other
     */
    function handleMerge(src, dst) {
      const srcRoot = find(src);
      const dstRoot = find(dst);
      if (subSets[srcRoot].order < subSets[dstRoot].order) {
        subSets[srcRoot].parent = dstRoot;
      } else if (subSets[srcRoot].order > subSets[dstRoot].order) {
        subSets[dstRoot].parent = srcRoot;
      } else {
        subSets[dstRoot].parent = srcRoot;
        subSets[srcRoot].order++;
      }
    }
  }
}

export function partTwo(input = null) {
  console.log({ day: 25, part: 2, value: "Merry Christmas" });
}

function buildGraph(input) {
  let graph = {};
  input
    .trim()
    .split("\n")
    .forEach((line) => {
      let [node, connections] = line.split(": ");
      connections = connections.split(" ");
      graph[node] = [...(graph[node] ?? []), ...connections];
      for (let connection of connections) {
        graph[connection] = [...(graph[connection] ?? []), node];
      }
    });
  return graph;
}
