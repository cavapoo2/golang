import * as React from "react";
import {EventListItem} from "./event_list_item";
import {Event} from "../model/event";

export interface EventListProps {
    userID:string;
    events: Event[];
    onEventBooked: (e: Event) => any;
}

export class EventList extends React.Component<EventListProps, {}> {
    public render() {
        console.log('EventList userid=',this.props.userID)
        const items = this.props.events.map(event =>
            <EventListItem key={event.ID} userID={this.props.userID} event={event} onBooked={() => this.props.onEventBooked(event)}/>
        );

        return <table className="table">
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Where</th>
                    <th colSpan={2}>When (start/end)</th>
                    <th>Actions</th>
                </tr>
            </thead>
            <tbody>
                {items}
            </tbody>
        </table>;
    }
}