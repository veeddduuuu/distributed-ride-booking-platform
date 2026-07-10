import './App.css';
import MapView from './components/MapView';
import { useTrip } from './hooks/useTrip';

const RIDES = [
  { id: 'suv', name: 'SUV', desc: 'Spacious ride for groups', icon: '🚐' },
  { id: 'sedan', name: 'Sedan', desc: 'Economic and comfortable', icon: '🚗' },
  { id: 'van', name: 'Van', desc: 'Perfect for larger groups', icon: '🚌' },
  { id: 'luxury', name: 'Luxury', desc: 'Premium experience', icon: '👑' },
];

function App() {
  const { pickup, destination, route, handleSetPickup } = useTrip();

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
          <h2>Select your desired ride</h2>
          
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
              <div key={ride.id} className="ride-item">
                <div className="ride-icon">{ride.icon}</div>
                <div className="ride-details">
                  <span className="ride-name">{ride.name}</span>
                  <span className="ride-desc">{ride.desc}</span>
                </div>
              </div>
            ))}
          </div>

          <button className="back-button" onClick={() => handleSetPickup({ latitude: 0, longitude: 0 }) /* Just a mock action */}>
            Back to Map
          </button>
        </div>
      </div>
    </div>
  );
}

export default App;
