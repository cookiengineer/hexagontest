/*
 * From https://www.redblobgames.com/grids/edges/
 * Copyright 2019 Red Blob Games <redblobgames@gmail.com>
 * License: Apache v2.0 <http://www.apache.org/licenses/LICENSE-2.0.html>
 */
'use strict';

console.info("I'm happy to answer questions about the code; email me at redblobgames@gmail.com");

const SPACING = 100;

function lerp(a, b, t) {
    return a * (1-t) + b * t;
}

function vlerp(a, b, t) {
    return {
        x: lerp(a.x, b.x, t),
        y: lerp(a.y, b.y, t),
        toString() { return `${this.x.toFixed(2)},${this.y.toFixed(2)}`; },
    };
}

/* Unified coordinate object:
   s === undefined: tile
   s === 'v': vertex northwest of tile
   s === 'W'|'N': edge to west or north of tile
*/
class Pos {
    constructor (q, r, s) {
        this.q = q;
        this.r = r;
        this.s = s;
    }

    get id() { return `${this.q},${this.r}${this.s ? ","+this.s : ""}`; }
    get x() { return (this.q + (this.s !== 'W' && this.s !== 'v'? 0.5 : 0.0)) * SPACING; }
    get y() { return (this.r + (this.s !== 'N' && this.s !== 'v'? 0.5 : 0.0)) * SPACING; }
    get xy() { return `${this.x.toFixed(2)},${this.y.toFixed(2)}`; }
    toString() { return this.id; }
    equals(other) { return this.id === other.id; }
}

class Tile extends Pos {
    get parity() { return (this.q+this.r) & 1; }
    
    /* Access functions from
       http://www-cs-students.stanford.edu/~amitp/game-programming/grids/
    */
    get neighbors() {
        let {q, r} = this;
        return [
            new Tile(q, r+1),
            new Tile(q+1, r),
            new Tile(q, r-1),
            new Tile(q-1, r),
        ];
    }

    get borders() {
        let {q, r} = this;
        return [
            new Edge(q, r, 'N'),
            new Edge(q, r, 'W'),
            new Edge(q, r+1, 'N'),
            new Edge(q+1, r, 'W'),
        ];
    }

    get corners() {
        let {q, r} = this;
        return [
            new Corner(q, r),
            new Corner(q, r+1),
            new Corner(q+1, r+1),
            new Corner(q+1, r),
        ];
    }

    makePolygon(blend) {
        return this.corners.map(p => vlerp(p, this, blend));
    }
}

class Edge extends Pos {
    get parity() { return this.s === 'N'? 0 : 1; }
    
    get joins() {
        let {q, r, s} = this;
        switch(s) {
        case 'N': return [new Tile(q, r-1), new Tile(q, r)];
        case 'W': return [new Tile(q-1, r), new Tile(q, r)];
        }
        throw "Invalid Edge";
    }

    get endpoints() {
        let {q, r, s} = this;
        switch(s) {
        case 'N': return [new Corner(q, r), new Corner(q+1, r)];
        case 'W': return [new Corner(q, r), new Corner(q, r+1)];
        }
        throw "Invalid Edge";
    }

    makePolygon(blend) {
        const {endpoints, joins} = this;
        return [
            endpoints[0].xy,
            vlerp(endpoints[0], joins[0], blend),
            vlerp(endpoints[1], joins[0], blend),
            endpoints[1].xy,
            vlerp(endpoints[1], joins[1], blend),
            vlerp(endpoints[0], joins[1], blend),
        ];
    }

    makeRectangle(blend) {
        const {endpoints, joins} = this;
        return [
            vlerp(endpoints[0], joins[0], blend),
            vlerp(endpoints[1], joins[0], blend),
            vlerp(endpoints[1], joins[1], blend),
            vlerp(endpoints[0], joins[1], blend),
        ];
    }
}

class Corner extends Pos {
    constructor(q, r) {
        super(q, r, 'v');
    }

    get touches() {
        let {q, r} = this;
        return [
            new Tile(q, r),
            new Tile(q, r-1),
            new Tile(q-1, r-1),
            new Tile(q-1, r),
        ];
    }
    
    makePolygon(blend) {
        return this.touches.map(p => vlerp(p, this, 1-blend));
    }
}


const shared = {
    computed: {
        SPACING() { return SPACING; },
        tiles() {
            let tiles = [];
            for (let q = 0; q < this.cols; q++) {
                for (let r = 0; r < this.rows; r++) {
                    tiles.push(new Tile(q, r));
                }
            }
            return tiles;
        },
        edges() {
            let edges = [];
            for (let q = 0; q <= this.cols; q++) {
                for (let r = 0; r <= this.rows; r++) {
                    if (q < this.cols) { edges.push(new Edge(q, r, 'N')); }
                    if (r < this.rows) { edges.push(new Edge(q, r, 'W')); }
                }
            }
            return edges;
        },
        corners() {
            let corners = [];
            for (let q = 0; q <= this.cols; q++) {
                for (let r = 0; r <= this.rows; r++) {
                    corners.push(new Corner(q, r));
                }
            }
            return corners;
        },
    },
    methods: {
        toggleTile(tile) {
            this.$set(this.tileState, tile, !this.tileState[tile]);
        },
        toggleEdge(edge) {
            this.$set(this.edgeState, edge, !this.edgeState[edge]);
        },
    },
};


new Vue({
    el: "#diagram-tile",
    mixins: [shared],
    data: {
        cols: 6,
        rows: 3,
        blend: 0,
    },
});

new Vue({
    el: "#diagram-edge",
    mixins: [shared],
    data: {
        cols: 6,
        rows: 3,
        blend: 1,
    },
});

new Vue({
    el: "#diagram-hybrid",
    mixins: [shared],
    data: {
        cols: 6,
        rows: 3,
        blend: 0.5,
    },
});

new Vue({
    el: "#diagram-corner-coordinates",
    mixins: [shared],
    data: { cols: 6, rows: 3, blend: 0.2, highlighted: {s:'?'}, },
});

new Vue({
    el: "#diagram-tile-wall",
    mixins: [shared],
    data: {
        cols: 9,
        rows: 5,
        tileState: {
            "0,1": true, "1,0": true, "2,0": true, "3,0": true, "4,0": true, "1,1": true, "1,2": true, "1,3": true, "1,4": true, "2,4": true, "3,4": true, "4,4": true, "5,4": true, "5,3": true, "5,1": true, "5,0": true,
        },
    },
});

new Vue({
    el: "#diagram-tile-pipe",
    mixins: [shared],
    data: {
        cols: 9,
        rows: 5,
        tileState: {
            "1,1": true,
            "2,1": true,
            "3,1": true,
            "4,1": true,
            "5,1": true,
            "6,1": true,
            "2,2": true,
            "2,3": true,
            "1,3": true,
            "5,2": true,
        },
    },
    computed: {
        edgeState() {
            // An edge is TRUE if both of its tiles are TRUE
            let state = {};
            for (let edge of this.edges) {
                state[edge] = edge.joins.every(t => this.tileState[t]);
            }
            return state;
        },
    },
});

new Vue({
    el: "#diagram-tile-light",
    mixins: [shared],
    data: {
        cols: 9,
        rows: 5,
        tileState: {
            "2,1": true, "2,2": true, "3,2": true, "3,1": true, "4,2": true, "3,3": true, "4,3": true, "5,2": true, "6,2": true, "1,2": true,
        },
    },
    computed: {
        edgeState() {
            // An edge is TRUE if either of its tiles are TRUE
            let state = {};
            for (let edge of this.edges) {
                state[edge] = edge.joins.some(t => this.tileState[t]);
            }
            return state;
        },
    },
});

new Vue({
    el: "#diagram-tile-thinwalls",
    mixins: [shared],
    data: {
        cols: 9,
        rows: 5,
        tileState: {
            "0,1": true, "1,0": true, "2,0": true, "3,0": true, "4,0": true, "1,1": true, "1,2": true, "1,3": true, "1,4": true, "2,4": true, "3,4": true, "4,4": true, "5,4": true, "5,3": true, "5,1": true, "5,0": true,
        },
    },
    computed: {
        edgeState() {
            // An edge is TRUE if either of its tiles are TRUE but not both
            let state = {};
            for (let edge of this.edges) {
                state[edge] = 1 === edge.joins.filter(t => this.tileState[t]).length;
            }
            return state;
        },
    },
});

new Vue({
    el: "#diagram-tile-coordinates",
    mixins: [shared],
    data: { cols: 9, rows: 5, },
});

new Vue({
    el: "#diagram-side-coordinates",
    mixins: [shared],
    data: { cols: 3, rows: 2, },
});

new Vue({
    el: "#diagram-edge-coordinates",
    mixins: [shared],
    data: { cols: 3, rows: 2, },
});

new Vue({
    el: "#diagram-borders",
    mixins: [shared],
    data: { cols: 4, rows: 3, highlighted: new Tile(2,1), },
});

new Vue({
    el: "#diagram-joins",
    mixins: [shared],
    data: { cols: 4, rows: 3, highlighted: new Edge(2,1,'N'), },
});


new Vue({
    el: "#diagram-edge-wall",
    mixins: [shared],
    data: {
        cols: 9,
        rows: 5,
        edgeState: { "0,2,N": true, "1,2,N": true, "2,1,N": true, "3,1,N": true, "5,1,W": true, "4,1,N": true, "2,1,W": true, "2,2,W": true, "2,3,W": true, "2,4,N": true, "3,4,N": true, "4,4,N": true, "5,3,W": true, },
    },
});

new Vue({
    el: "#diagram-edge-pipe",
    mixins: [shared],
    data: {
        cols: 9,
        rows: 5,
        edgeState: {
            "1,1,N": true,
            "1,2,W": true,
            "2,1,W": true,
            "2,2,W": true,
            "3,1,W": true,
            "3,2,W": true,
            "4,2,W": true,
            "5,2,W": true,
            "4,3,W": true,
            "3,3,W": true,
            "3,3,N": true,
        },
    },
    computed: {
        tileState() {
            // A tile is TRUE if either of its edges are TRUE
            let state = {};
            for (let tile of this.tiles) {
                state[tile] = tile.borders.some(e => this.edgeState[e]);
            }
            return state;
        },
    },
});

