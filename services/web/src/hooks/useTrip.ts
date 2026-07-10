import {useState, useEffect} from 'react';
import type { Coordinates, Route } from './../api/trip';
import { fetchTripPreview } from './../api/trip';

export function useTrip(){
    const [pickup, setPickup] = useState<Coordinates | null>(null);
    const [destination, setDestination] = useState<Coordinates | null>(null);
    const [route, setRoute] = useState<Route | null>(null);

    const handleSetPickup = (coords: Coordinates) => {
        if(!pickup){
            setPickup(coords);
        }
        else if(!destination){
            setDestination(coords);
        }
        else{
            setPickup(coords);
            setDestination(null);
            setRoute(null);
        }
    };

    useEffect(() => {
        if(pickup && destination){
            fetchTripPreview(pickup, destination)
                .then((route) => {
                    setRoute(route);
                })
                .catch((error) => {
                    console.error('Error fetching trip preview:', error);
                });
        }
    }, [pickup, destination]);

    return {
        pickup,
        destination,
        route,
        handleSetPickup,
    };
}

