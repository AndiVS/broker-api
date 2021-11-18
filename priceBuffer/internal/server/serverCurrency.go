package server

import (
	"priceBuffer/protocol"
)

// CurrencyServer for grpc
type CurrencyServer struct {
	*protocol.UnimplementedCurrencyServiceServer
}

/*func (s *CurrencyServer) GetPrice(stream *protocol.GetPriceRequest) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		key := serialize(in.Location)

		s.mu.Lock()
		s.routeNotes[key] = append(s.routeNotes[key], in)
		// Note: this copy prevents blocking other clients while serving this one.
		// We don't need to do a deep copy, because elements in the slice are
		// insert-only and never modified.
		rn := make([]*pb.RouteNote, len(s.routeNotes[key]))
		copy(rn, s.routeNotes[key])
		s.mu.Unlock()

		for _, note := range rn {
			if err := stream.Send(note); err != nil {
				return err
			}
		}
	}
}


func (s *CurrencyServer) Stream(r *protocol.GetPriceRequest, srv Streamer_StreamServer) error {
	client := s.pubsub.NewClient()
	defer client.Close()

	// use request id as unique topic name
	if err := client.Subscribe(r.Id); err != nil {
		return err
	}

	// start the background producer
	_, err := s.queue.Push(&jobs.Job{
		Job:     "app.job.Produce",
		Payload: `{"requestID":"` + r.Id + `"}`,
		Options: &jobs.Options{},
	})
	if err != nil {
		return err
	}

	// forward data from topic to stream
	for msgData := range client.Channel() {
		msg := &Message{}
		if err := json.Unmarshal(msgData.Payload, msg); err != nil {
			return err
		}

		if err := srv.Send(&Data{
			Sequence: int32(msg.SequenceID),
			Data:     []byte(msg.Data),
		}); err != nil {
			return err
		}
	}

	return nil
}*/