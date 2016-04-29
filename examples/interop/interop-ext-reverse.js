var start = new Date().getTime()

setInterval(function () {
    ReactDOM.render(
        Elapser({
            elapsed: Math.round(((new Date().getTime() - start) / 1000)),
            title: "Interop from JavaScript"
        }),
        document.getElementById('react-reverse')
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
