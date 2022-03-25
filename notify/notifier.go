package notify

type Notifier interface {
	// Notify(to string, campground, checkInDate, checkOutDate string, available []string) error
	Notify(to string, available []string) error
}
