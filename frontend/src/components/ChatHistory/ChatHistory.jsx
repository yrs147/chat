import React from "react";
import "./ChatHistory.scss"


class ChatHistory extends Component {
    render() {
        const messages = this.props.ChatHistory.map((msg, index) =>(
            <p key={index}>{msg.data}</p>
        ));

        return (
            <div className="ChatHistory">
                <h2>Chat History</h2>
                {messages}
            </div>
        );
    }
}

export default ChatHistory;