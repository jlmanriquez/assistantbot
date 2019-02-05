const express = require('express');
const app = express();

app.use(express.static('public'));

app.use('/css', express.static(__dirname + '/public/css'));
app.use('/js', express.static(__dirname + '/public/js'));
app.use('/images', express.static(__dirname + '/public/images'));

const server = app.listen(8082, () => {
    const port = server.address().port;
    console.log('Server iniciado en puerto 8082');
});