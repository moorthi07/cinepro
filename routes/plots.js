const express = require('express');
const router = express.Router();

// Replace with your actual data store
const plots = [
    { id: 1, title: 'Plot A', description: 'Description A', cineProjectId: 1 },
    { id: 2, title: 'Plot B', description: 'Description B', cineProjectId: 2 }
];

// GET /plots
router.get('/', (req, res) => {
    res.json(plots);
});

// POST /plots
router.post('/', (req, res) => {
    const newPlot = req.body;
    plots.push(newPlot);
    res.status(201).json(newPlot);
});

// GET /plots/:plotId
router.get('/:plotId', (req, res) => {
    const plotId = parseInt(req.params.plotId);
    const plot = plots.find(p => p.id === plotId);
    if (plot) {
        res.json(plot);
    } else {
        res.status(404).json({ error: 'Plot not found' });
    }
});

// PUT /plots/:plotId
router.put('/:plotId', (req, res) => {
    const plotId = parseInt(req.params.plotId);
    const updatedPlot = req.body;
    const index = plots.findIndex(p => p.id === plotId);
    if (index !== -1) {
        plots[index] = updatedPlot;
        res.json(updatedPlot);
    } else {
        res.status(404).json({ error: 'Plot not found' });
    }
});

// DELETE /plots/:plotId
router.delete('/:plotId', (req, res) => {
    const plotId = parseInt(req.params.plotId);
    const index = plots.findIndex(p => p.id === plotId);
    if (index !== -1) {
        plots.splice(index, 1);
        res.sendStatus(204);
    } else {
        res.status(404).json({ error: 'Plot not found' });
    }
});

module.exports = router;