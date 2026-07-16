import type { RideShare } from '../api/trip';

interface FareCardProps {
  fare: RideShare;
  selected: boolean;
  onClick: (fare: RideShare) => void;
}

const FARE_DETAILS: Record<string, { name: string; desc: string; icon: string }> = {
  sedan: { name: 'Sedan', desc: 'Economic and comfortable', icon: '🚗' },
  suv: { name: 'SUV', desc: 'Spacious ride for groups', icon: '🚐' },
  van: { name: 'Van', desc: 'Perfect for larger groups', icon: '🚌' },
  luxury: { name: 'Luxury', desc: 'Premium experience', icon: '👑' },
};

export default function FareCard({ fare, selected, onClick }: FareCardProps) {
  const details = FARE_DETAILS[fare.packageSlug] || {
    name: fare.packageSlug,
    desc: 'Standard ride',
    icon: '🚙',
  };

  return (
    <div
      className={`fare-card ${selected ? 'fare-card--selected' : ''}`}
      onClick={() => onClick(fare)}
      role="button"
      tabIndex={0}
      onKeyDown={(e) => {
        if (e.key === 'Enter' || e.key === ' ') {
          e.preventDefault();
          onClick(fare);
        }
      }}
    >
      <div className="fare-card-icon">{details.icon}</div>
      <div className="fare-card-details">
        <span className="fare-card-name">{details.name}</span>
        <span className="fare-card-desc">{details.desc}</span>
      </div>
      <div className="fare-card-price">
        ₹ {fare.totalPrice.toFixed(2)}
      </div>
    </div>
  );
}
