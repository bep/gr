
// This is the one exported component from the Go side.
var Elapser = require('./interop').Elapser;

var start = new Date().getTime()

// Use a component set on the Go side as a global variable.
setInterval(function () {
    ReactDOM.render(
        ElapserGlobal({
            elapsed: Math.round(((new Date().getTime() - start) / 1000)),
            title: "Interop from JavaScript Global"
        }),
        document.getElementById('react-reverse-global')
    );
}, 50);


setInterval(function () {
    ReactDOM.render(
        Elapser({
            elapsed: Math.round(((new Date().getTime() - start) / 1000)),
            title: "Interop from JavaScript Require"
        }),
        document.getElementById('react-reverse-require')
    );
}, 50);


var Hello = React.createClass({
    render: function () {
        console.log("JS: render")
        return React.DOM.h1(null, "JS Render: Go React!")
    },

    componentDidMount: function () {
        console.log("JS: componentDidMount")
    }
});


ReactDOM.render(
    React.createElement(Hello),
    document.getElementById('react-hello')
);
