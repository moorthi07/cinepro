const express = require('express');
const bodyParser = require('body-parser');

const plotRoutes = require('./routes/plots');
const screenplayRoutes = require('./routes/screenplay');


const app = express();
const port = 3000;

// Replace with your actual data store
const cineProjects = [
    { id: 1, title: 'Project A', description: 'Description A' },
    { id: 2, title: 'Project B', description: 'Description B' }
];

app.use(bodyParser.json());

app.use('/plots', plotRoutes);
app.use('/screenplays', screenplayRoutes);

// GET /cineProjects
app.get('/cineProjects', (req, res) => {
    console.log('cineProjects')
    res.json(cineProjects);
});

// POST /cineProjects
app.post('/cineProjects', (req, res) => {
    const newProject = req.body;
    cineProjects.push(newProject);
    res.status(201).json(newProject);
});

// GET /cineProjects/:cineProjectId
app.get('/cineProjects/:cineProjectId', (req, res) => {
    const cineProjectId = parseInt(req.params.cineProjectId);
    const project = cineProjects.find(p => p.id === cineProjectId);
    if (project) {
        res.json(project);
    } else {
        res.status(404).json({ error: 'Cine project not found' });
    }
});

// PUT /cineProjects/:cineProjectId
app.put('/cineProjects/:cineProjectId', (req, res) => {
    const cineProjectId = parseInt(req.params.cineProjectId);
    const updatedProject = req.body;
    const index = cineProjects.findIndex(p => p.id === cineProjectId);
    if (index !== -1) {
        cineProjects[index] = updatedProject;
        res.json(updatedProject);
    } else {
        res.status(404).json({ error: 'Cine project not found' });
    }
});

// DELETE /cineProjects/:cineProjectId
app.delete('/cineProjects/:cineProjectId', (req, res) => {
    const cineProjectId = parseInt(req.params.cineProjectId);
    const index = cineProjects.findIndex(p => p.id === cineProjectId);
    if (index !== -1) {
        cineProjects.splice(index, 1);
        res.sendStatus(204);
    } else {
        res.status(404).json({ error: 'Cine project not found' });
    }
});

app.listen(port, () => {
    console.log(`Server listening on port ${port}`);
});