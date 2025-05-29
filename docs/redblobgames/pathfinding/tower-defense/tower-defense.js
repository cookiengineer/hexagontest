// From https://www.redblobgames.com/pathfinding/tower-defense/
// Copyright 2014 Red Blob Games <redblobgames@gmail.com>
// License: Apache v2.0 <http://www.apache.org/licenses/LICENSE-2.0.html>

import {SearchOptions} from '../search.js';
import {SquareGrid, SquareGridLayout} from '../grid.js';
import {Diagram, Layers, Markers, addSliderSteps, clamp, lerp} from '../diagram.js';

console.info("I'm happy to answer questions about the code; email me at redblobgames@gmail.com");

const parentPointerStyle = {
    fillStyle: "hsl(0, 0%, 100%)",
    strokeStyle: "hsla(0, 0%, 0%, 0.5)",
    lineWidth: 1.0,
};
const COLORS = {
    units: "hsl(0, 20%, 50%)",
    current: "hsl(220, 60%, 60%)",
};


/* Paint using a gradient */
function paintGradient(field, scale) {
    const hues = [350, 400], saturations = [40, 10], lightnesses = [65, 80];
    return {
        override(ctx, ss, map, id) {
            if (map[id].tile_weight === 1) {
                let t = clamp(map[id][field] / scale, 0, 1);
                let hue = lerp(hues[0], hues[1], t),
                    saturation = lerp(saturations[0], saturations[1], t),
                    lightness = lerp(lightnesses[0], lightnesses[1], t);
                ctx.fillStyle = `hsl(${hue|0},${saturation|0}%,${lightness|0}%)`;
            }
        },
    };
}


class Particle {
    constructor (graph) {
        this.phase = Math.random();
        this.reset(graph);
    }

    reset(graph) {
        this.id = Math.floor(Math.random() * graph.num_nodes);
        this.dx = 0.8 * (Math.random() - 0.5);
        this.dy = 0.8 * (Math.random() - 0.5);
    }

    update(graph, map, dtSeconds) {
        this.phase += dtSeconds;
        if (this.phase >= 1.0) {
            this.phase %= 1.0;
            this.id = map[this.id].parent;
            if (this.id === undefined) {
                this.reset(graph);
            }
        }
    }

    position(map, layout) {
        let next = map[this.id].parent;
        if (map[this.id].tile_weight === Infinity) { return [-999, -999]; } // hide
        if (next === undefined) { next = this.id; }

        let xy0 = layout.graph.from_id(this.id);
        let xy1 = layout.graph.from_id(next);
        let x = this.dx + xy0[0] * (1.0-this.phase) + xy1[0] * this.phase;
        let y = this.dy + xy0[1] * (1.0-this.phase) + xy1[1] * this.phase;
        return layout.xy_scaled([x, y]);
    }
}


function unitLayer(graph, buttonEl) {
    const speed = 3.0;
    let particles = [];
    let animation_id = 0;
    let lastTimestamp = Date.now();
    let map = [];
    let diagram = null;

    buttonEl = document.querySelector(buttonEl);

    for (let i = 0; i < 300; i++) {
        particles.push(new Particle(graph));
    }

    function draw(ctx, ss, map_) {
        diagram = this;
        map = map_;

        ctx.fillStyle = COLORS.units;
        for (let p of particles) {
            let [x, y] = p.position(map, this.layout);
            ctx.beginPath();
            ctx.arc(x, y, 3, 0, 2*Math.PI);
            ctx.fill();
        }
    }

    function update(p) {
        let now = Date.now();
        let dtMilliseconds = Math.min(500, now - lastTimestamp);
        if (diagram && dtMilliseconds > 1000/30) {
            lastTimestamp = now;
            particles.forEach(p => p.update(graph, map, speed * dtMilliseconds/1000));
            diagram.redraw();
        }
        if (animation_id) {
            animation_id = requestAnimationFrame(update);
        }
    }

    function togglePlayPause() {
        if (animation_id) {
            cancelAnimationFrame(animation_id);
            animation_id = 0;
            buttonEl.textContent = "Animate";
        } else {
            lastTimestamp = Date.now();
            animation_id = requestAnimationFrame(update);
            buttonEl.textContent = "Pause";
        }
    };
    buttonEl.addEventListener('click', togglePlayPause);

    return {draw, togglePlayPause};
}

let shared_graph = new SquareGrid(20, 10);
let graph_layout = new SquareGridLayout(shared_graph, 30);
[10, 15, 24, 30, 44, 50, 55, 64, 75, 80, 82, 83, 84, 85, 86, 90, 95, 102, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 122, 134, 138, 142, 144, 154, 156, 158, 162, 164, 176, 178, 184, 194, 196].forEach((id) => shared_graph.set_tile_weight(id, Infinity));

let frontier_options = new SearchOptions([shared_graph.to_id(6, 8)]);
let frontier_diagram = new Diagram(
    "#diagram-frontier", shared_graph, frontier_options, graph_layout,
    Layers.base({frontier: {}}),
    Layers.frontierTiles(),
    Layers.editable(1, Infinity),
    Layers.draggableMarker(Markers.blob(15), frontier_options.starts, 0)
);
addSliderSteps("#diagram-frontier", [frontier_diagram]).setSliderPosition(2);
    
let expansion_options = new SearchOptions([shared_graph.to_id(6, 8)]);
let expansion_diagram = new Diagram(
    "#diagram-expansion", shared_graph, expansion_options, graph_layout,
    Layers.base({current: {fillStyle: COLORS.current},}),
    Layers.frontierTiles(),
    Layers.neighborTiles(),
    Layers.editable(1, Infinity),
    Layers.draggableMarker(Markers.blob(15), expansion_options.starts, 0)
);
addSliderSteps("#diagram-expansion", [expansion_diagram]).setSliderPosition(2);


let distances_options = new SearchOptions([shared_graph.to_id(6, 8)]);
let distances_diagram = new Diagram(
    "#diagram-distances", shared_graph, distances_options, graph_layout,
    Layers.base(),
    Layers.frontierTiles(),
    Layers.numericLabels('cost_so_far'),
    Layers.editable(1, Infinity),
    Layers.draggableMarker(Markers.blob(15), distances_options.starts, 0)
);
addSliderSteps("#diagram-distances", [distances_diagram]).setSliderPosition(10);


let parents_options = new SearchOptions([shared_graph.to_id(6, 8)]);
let parents_diagram = new Diagram(
    "#diagram-parents", shared_graph, parents_options, graph_layout,
    Layers.base(),
    Layers.frontierTiles(),
    Layers.parentPointers(false, 3, parentPointerStyle),
    Layers.editable(1, Infinity),
    Layers.draggableMarker(Markers.blob(15), parents_options.starts, 0)
);
addSliderSteps("#diagram-parents", [parents_diagram]).setSliderPosition(10);


let units_options = new SearchOptions([shared_graph.to_id(6, 8)]);
let units_layer = unitLayer(shared_graph, "#unit-flow-button");
let units_diagram = new Diagram(
    "#diagram-units", shared_graph, units_options, graph_layout,
    Layers.base(),
    Layers.editable(1, Infinity),
    units_layer,
    Layers.draggableMarker(Markers.blob(15), units_options.starts, 0),
);
units_layer.togglePlayPause();
units_layer.togglePlayPause();


let concepts_options = new SearchOptions([shared_graph.to_id(6, 8)]);
let concepts_diagram = new Diagram(
    "#diagram-concepts", shared_graph, concepts_options, graph_layout,
    Layers.base(paintGradient('cost_so_far', 20)),
    Layers.parentPointers(false, 3, {fillStyle: "hsla(30, 0%, 0%, 0.1)"}),
    Layers.numericLabels('cost_so_far', "hsl(30, 10%, 90%)"),
    Layers.editable(1, Infinity),
    Layers.draggableMarker(Markers.blob(15), concepts_options.starts, 0)
);
