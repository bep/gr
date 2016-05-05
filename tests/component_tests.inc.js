
global.Hello = React.createClass({
    render: function () {
		var message = this.props.message
        return React.DOM.h1(null, message)
    }

});
