import { useState } from 'react';
import '../App.css';
import MapView from '../components/MapView';
import { useTrip } from '../hooks/useTrip';
import { useRiderWebSocket } from '../hooks/useRiderWebSocket';
import { Link } from 'react-router-dom';

const RIDES = [
  { id: 'suv', name: 'SUV', desc: 'Spacious ride for groups', icon: '🚐' },
  { id: 'sedan', name: 'Sedan', desc: 'Economic and comfortable', icon: '🚗' },
  { id: 'van', name: 'Van', desc: 'Perfect for larger groups', icon: '🚌' },
  { id: 'luxury', name: 'Luxury', desc: 'Premium experience', icon: '👑' },
];

// Generate a random ID for this session
const RIDER_ID = 'rider_' + Math.random().toString(36).substr(2, 9);

export function RiderPage() {
  const { pickup, destination, route, handleSetPickup } = useTrip();
  const { isConnected, sendRideRequest } = useRiderWebSocket(RIDER_ID);
  const [selectedRide, setSelectedRide] = useState<string | null>(null);

  const handleRequestRide = () => {
    if (pickup && destination && selectedRide) {
      sendRideRequest({
        pickup,
        destination,
        packageSlug: selectedRide,
      });
      alert('Ride request sent to backend via WebSocket!');
    } else {
      alert('Please select pickup, destination, and ride type.');
    }
  };

  return (
    <div className="app-container">
      <div className="map-container">
        <MapView 
          pickup={pickup} 
          destination={destination} 
          route={route} 
          onMapClick={(lat, lng) => handleSetPickup({ latitude: lat, longitude: lng })} 
        />
      </div>
      
      <div className="sidebar">
        <div className="card">
          <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
            <h2>Select your desired ride</h2>
            <div title="WebSocket Status">
               {isConnected ? '🟢 Connected' : '🔴 Disconnected'}
            </div>
          </div>
          
          <div className="route-info">
            <p className="routing-text">
              {route ? `Routing for ${(route.distance / 1000).toFixed(2)} km` : 'Select pickup and destination on the map'}
            </p>
            {route && (
              <p className="time-text">
                <span className="clock-icon">🕒</span> You'll arrive in: {Math.round(route.duration / 60)} minutes
              </p>
            )}
          </div>

          <div className="ride-list">
            {RIDES.map((ride) => (
              <div 
                key={ride.id} 
                className={`ride-item ${selectedRide === ride.id ? 'selected' : ''}`}
                onClick={() => setSelectedRide(ride.id)}
                style={{
                  cursor: 'pointer',
                  border: selectedRide === ride.id ? '2px solid #000' : '2px solid transparent',
                  borderRadius: '12px',
                  padding: '12px'
                }}
              >
                <div className="ride-icon">{ride.icon}</div>
                <div className="ride-details">
                  <span className="ride-name">{ride.name}</span>
                  <span className="ride-desc">{ride.desc}</span>
                </div>
              </div>
            ))}
          </div>

          <div style={{ display: 'flex', gap: '1rem', marginTop: '1rem' }}>
            <button 
              className="back-button" 
              style={{ flex: 1, backgroundColor: '#000', color: '#fff' }}
              onClick={handleRequestRide}
              disabled={!pickup || !destination || !selectedRide}
            >
              Request Ride
            </button>
            <Link to="/" className="back-button" style={{ flex: 1, textDecoration: 'none', textAlign: 'center' }}>
              Home
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
}
