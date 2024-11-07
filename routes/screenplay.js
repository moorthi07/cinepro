const express = require('express');
const router = express.Router();

// Replace with your actual data store
const screenplays = [
    { id: 1, title: 'Screenplay A', content: 'Content A', cineProjectId: 1 },
    { id: 2, title: 'Screenplay B', content: 'Content B', cineProjectId: 2 }
];

// GET /screenplays
router.get('/', (req, res) => {
    res.json(screenplays);
});

// POST /screenplays
router.post('/', (req, res) => {
    const newScreenplay = req.body;
    screenplays.push(newScreenplay);
    res.status(201).json(newScreenplay);
});

// GET /screenplays/:screenplayId
router.get('/:screenplayId', (req, res) => {
    const screenplayId = parseInt(req.params.screenplayId);
    const screenplay = screenplays.find(s => s.id === screenplayId);
    if (screenplay) {
        res.json(screenplay);
    } else {
        res.status(404).json({ error: 'Screenplay not found' });
    }
});

// PUT /screenplays/:screenplayId
router.put('/:screenplayId', (req, res) => {
    const screenplayId = parseInt(req.params.screenplayId);
    const updatedScreenplay = req.body;
    const index = screenplays.findIndex(s => s.id === screenplayId);
    if (index !== -1) {
        screenplays[index] = updatedScreenplay;
        res.json(updatedScreenplay);
    } else {
        res.status(404).json({ error: 'Screenplay not found' });
    }
});

// DELETE /screenplays/:screenplayId
router.delete('/:screenplayId', (req, res) => {
    const screenplayId = parseInt(req.params.screenplayId);
    const index = screenplays.findIndex(s => s.id === screenplayId);
    if (index !== -1) {
        screenplays.splice(index, 1);
        res.sendStatus(204);
    } else {
        res.status(404).json({ error: 'Screenplay not found' });
    }
});

module.exports = router;