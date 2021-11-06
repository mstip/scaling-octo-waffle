import * as d3 from "https://cdn.skypack.dev/d3@6";
import * as d3Hierarchy from "https://cdn.skypack.dev/d3-hierarchy@3";
import * as d3Zoom from "https://cdn.skypack.dev/d3-zoom@3";



function graph(root, {
  label = d => d.data.name,
  highlight = () => false,
  marginLeft = 250,
  marginTop = 20
} = {}) {
  const dy = 120
  const dx = 12

  const tree = d3.tree().nodeSize([50, 50])
  const treeLink = d3.linkHorizontal().x(d => d.x).y(d => d.y)
  root = tree(root);

  let x0 = Infinity;
  let x1 = -x0;
  root.each(d => {
    if (d.x > x1) x1 = d.x;
    if (d.x < x0) x0 = d.x;
  });

  const zoom = d3Zoom.zoom()
  .scaleExtent([1, 8])
  .on("zoom", event => {
    const {transform} = event;
    g.attr("transform", transform);
    g.attr("stroke-width", 1 / transform.k);
  });

  const svg = d3.create("svg")
    .attr("viewBox", [0, 0, 500, x1 - x0 + dx * 2])
    .style("overflow", "visible")
    .call(zoom)
    // .call(d3.zoom().on("zoom", el => {
      // svg.attr("transform", d3.event.transform)
  //  }))

  const g = svg.append("g")
    .attr("font-family", "sans-serif")
    .attr("font-size", 8)
    .attr("transform", `translate(${marginLeft},${marginTop})`);

  const link = g.append("g")
    .attr("fill", "none")
    .attr("stroke", "#555")
    .attr("stroke-opacity", 0.4)
    .attr("stroke-width", 1.5)
    .selectAll("path")
    .data(root.links())
    .join("path")
    .attr("stroke", d => highlight(d.source) && highlight(d.target) ? "red" : null)
    .attr("stroke-opacity", d => highlight(d.source) && highlight(d.target) ? 1 : null)
    .attr("d", treeLink);

  const node = g.append("g")
    .attr("stroke-linejoin", "round")
    .attr("stroke-width", 3)
    .selectAll("g")
    .data(root.descendants())
    .join("g")
    .attr("transform", d => `translate(${d.x},${d.y})`)
    .attr('id', 'marcsid')
    .on('click', event => console.log(event))
    .on('mouseover', el => el.srcElement.setAttribute('fill','green'))

  // node.append("circle")
  //   .attr("fill", d => highlight(d) ? "red" : d.children ? "#555" : "#999")
  //   .attr("r", 2.5);

  node.append("path")
  .attr('d', 'M3 2.5a2.5 2.5 0 0 1 5 0 2.5 2.5 0 0 1 5 0v.006c0 .07 0 .27-.038.494H15a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1v7.5a1.5 1.5 0 0 1-1.5 1.5h-11A1.5 1.5 0 0 1 1 14.5V7a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h2.038A2.968 2.968 0 0 1 3 2.506V2.5zm1.068.5H7v-.5a1.5 1.5 0 1 0-3 0c0 .085.002.274.045.43a.522.522 0 0 0 .023.07zM9 3h2.932a.56.56 0 0 0 .023-.07c.043-.156.045-.345.045-.43a1.5 1.5 0 0 0-3 0V3zM1 4v2h6V4H1zm8 0v2h6V4H9zm5 3H9v8h4.5a.5.5 0 0 0 .5-.5V7zm-7 8V7H2v7.5a.5.5 0 0 0 .5.5H7z')
  .attr('fill', 'black')
  .attr("transform", `scale(0.5)`);
  

  node.append("text")
    .attr("fill", d => highlight(d) ? "red" : null)
    .attr("dy", "0.31em")
    .attr("x", d => d.children ? -6 : 6)
    .attr("text-anchor", d => d.children ? "end" : "start")
    .text(label)
    .clone(true).lower()
    .attr("stroke", "white");

  return svg.node();
}



const tabledata = d3.stratify()
  .id(function (d) { return d.name; })
  .parentId(function (d) { return d.parent; })
  ([
    { "name": "Eve", "parent": "" },
    { "name": "Cain", "parent": "Eve" },
    { "name": "Seth", "parent": "Eve" },
    { "name": "Enos", "parent": "Seth" },
    { "name": "Noam", "parent": "Seth" },
    { "name": "Abel", "parent": "Eve" },
    { "name": "Awan", "parent": "Eve" },
    { "name": "Enoch", "parent": "Awan" },
    { "name": "Azura", "parent": "Eve" }
  ]);

const treedata = d3.hierarchy({
  name: "root",
  children: [
    { name: "child #1" },
    {
      name: "child #2",
      children: [
        { name: "grandchild #1" },
        { name: "grandchild #2" },
        { name: "grandchild #3" }
      ]
    }
  ]
});


document.body.appendChild(graph(tabledata));