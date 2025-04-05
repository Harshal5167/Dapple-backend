package service

// import (
// 	"encoding/json"
// 	"fmt"

// 	"github.com/Harshal5167/Dapple-backend/internal/dto/request"
// 	"github.com/Harshal5167/Dapple-backend/internal/interfaces"
// 	"github.com/gofiber/contrib/socketio"
// )

// type SocketService struct {
// 	testService interfaces.TestService
// }

// func NewSocketService(ts interfaces.TestService) *SocketService {
// 	return &SocketService{
// 		testService: ts,
// 	}
// }

// func (s *SocketService) HandleConnection(kws *socketio.Websocket) {
// 	socketio.On(socketio.EventConnect, s.handleConnect)
// 	socketio.On(socketio.EventMessage, s.handleMessage)
// 	socketio.On(socketio.EventDisconnect, s.handleDisconnect)
// 	socketio.On(socketio.EventClose, s.handleClose)
// }

// func (s *SocketService) handleConnect(ep *socketio.EventPayload) {
// 	fmt.Println("Connected to socket")
// 	ep.Kws.Emit([]byte(`{"status": "success", "message": "Connected"}`), socketio.TextMessage)
// }

// func (s *SocketService) handleMessage(ep *socketio.EventPayload) {
// 	var message = &request.TestData{}

// 	err := json.Unmarshal(ep.Data, &message)
// 	if err != nil {
// 		ep.Kws.Emit([]byte(`{"status": "error", "message": "Invalid JSON"}`), socketio.TextMessage)
// 		return
// 	}

// 	fmt.Println("Received message: ", message.Event, message.SessionId, message.QuestionId)
// 	if message.Event == string(request.Image) {
// 		ep.Kws.Emit([]byte(`{"status": "success", "message": "Data received"}`), socketio.TextMessage)
// 		err = s.testService.EvaluateImageAnswer(message)
// 		if err != nil {
// 			ep.Kws.Emit([]byte(`{"status": "error", "message": "Error evaluating image"}`), socketio.TextMessage)
// 			return
// 		}
// 	}

// 	if message.Event == string(request.Text) {
// 		ep.Kws.Emit([]byte(`{"status": "success", "message": "Data received"}`), socketio.TextMessage)
// 		isEnd, err := s.testService.EvaluateTestAnswer(message)
// 		if err != nil {
// 			ep.Kws.Emit([]byte(`{"status": "error", "message": "Error evaluating text"}`), socketio.TextMessage)
// 			return
// 		}
// 		if isEnd {
// 			ep.Kws.Emit([]byte(`{"status": "success", "message": "Connection Closed"}`), socketio.TextMessage)
// 			ep.Kws.Close()
// 			return
// 		}
// 	}

// 	if message.Event == string(request.Retry) {
// 		err = s.testService.RetryQuestion(message.SessionId, message.QuestionId)
// 		if err != nil {
// 			ep.Kws.Emit([]byte(`{"status": "error", "message": "Error clearing question frames"}`), socketio.TextMessage)
// 			return
// 		}
// 		ep.Kws.Emit([]byte(`{"status": "success", "message": "cleared question frames"}`), socketio.TextMessage)
// 	}
// }

// func (s *SocketService) handleDisconnect(ep *socketio.EventPayload) {
// 	fmt.Println("Disconnected from socket")
// 	ep.Kws.Emit([]byte(`{"status": "success", "message": "Disconnected"}`))
// }

// func (s *SocketService) handleClose(ep *socketio.EventPayload) {
// 	fmt.Println("Closed socket connection")
// 	ep.Kws.Emit([]byte(`{"status": "success", "message": "Connection Closed"}`))
// }
