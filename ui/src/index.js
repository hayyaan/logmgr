import React from 'react';
import ReactDOM from 'react-dom';

export default
class App extends React.Component {

    componentDidMount() {
        if (!!window.EventSource) {
            var source = new EventSource('/logs');

            source.addEventListener('message', function (e) {
                console.log(e.data);
            }, false);

            source.addEventListener('open', function (e) {
                console.log('listening for new logs');
            }, false);

            source.addEventListener('error', function (e) {
                if (e.readyState == EventSource.CLOSED) {
                    alert("Connection closed");
                }
            }, false);
        } else {
            alert("Upgrade your browser, WTF");
        }
    }

    render() {
        return (
            <div>Hello Logs</div>
        );
    }
}

ReactDOM.render(<App/>, document.querySelector('#app'));
