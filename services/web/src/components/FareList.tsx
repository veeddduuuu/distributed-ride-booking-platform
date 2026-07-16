import type { RideShare } from '../api/trip';
import FareCard from './FareCard';

interface FareListProps {
  fares: RideShare[];
  selectedFare: RideShare | null;
  onSelectFare: (fare: RideShare) => void;
}

export default function FareList({ fares, selectedFare, onSelectFare }: FareListProps) {
  if (!fares || fares.length === 0) {
    return null;
  }

  return (
    <div className="fare-list">
      {fares.map((fare) => (
        <FareCard
          key={`${fare.packageSlug}-${fare.totalPrice}`}
          fare={fare}
          selected={selectedFare?.packageSlug === fare.packageSlug}
          onClick={onSelectFare}
        />
      ))}
    </div>
  );
}
