package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"os"

	pokemonpc "github.com/TRomesh/grpc-pokemon/pokemon"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

const defaultPort = "4041"

func main() {

	fmt.Println("Pokemon Client")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:4041", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close() // Может, в отдельной функции ошибка будет обработана?

	c := pokemonpc.NewPokemonServiceClient(cc)

	// создаем покемона
	fmt.Println("Creating the pokemon")
	pokemon := &pokemonpc.Pokemon{
		Pid:         "Poke01",
		Name:        "Pikachu",
		Power:       "Fire",
		Description: "Fluffy",
	}
	createPokemonRes, err := c.CreatePokemon(context.Background(), &pokemonpc.CreatePokemonRequest{Pokemon: pokemon})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Pokemon has been created: %v", createPokemonRes)
	pokemonID := createPokemonRes.GetPokemon().GetId()

	// считываем покемона
	fmt.Println("Reading the pokemon")
	readPokemonReq := &pokemonpc.ReadPokemonRequest{Pid: pokemonID}
	readPokemonRes, readPokemonErr := c.ReadPokemon(context.Background(), readPokemonReq)
	if readPokemonErr != nil {
		fmt.Printf("Error happened while reading: %v \n", readPokemonErr)
	}

	fmt.Printf("Pokemon was read: %v \n", readPokemonRes)

	// обновляем покемона
	newPokemon := &pokemonpc.Pokemon{
		Id:          pokemonID,
		Pid:         "Poke01",
		Name:        "Pikachu",
		Power:       "Fire Fire Fire",
		Description: "Fluffy",
	}
	updateRes, updateErr := c.UpdatePokemon(context.Background(), &pokemonpc.UpdatePokemonRequest{Pokemon: newPokemon})
	if updateErr != nil {
		fmt.Printf("Error happened while updating: %v \n", updateErr)
	}
	fmt.Printf("Pokemon was updated: %v\n", updateRes)

	// удаляем покемона
	deleteRes, deleteErr := c.DeletePokemon(context.Background(), &pokemonpc.DeletePokemonRequest{Pid: pokemonID})

	if deleteErr != nil {
		fmt.Printf("Error happened while deleting: %v \n", deleteErr)
	}
	fmt.Printf("Pokemon was deleted: %v \n", deleteRes)

	// выводим список покемонов

	stream, err := c.ListPokemon(context.Background(), &pokemonpc.ListPokemonRequest{})
	if err != nil {
		log.Fatalf("error while calling ListPokemon RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetPokemon())
	}
}

var collection *mongo.Collection

type server struct {
	pokemonpc.PokemonServiceServer
}
type pokemonItem struct {
	Id          primitive.ObjectID
	Pid         string
	Name        string
	Power       string
	Description string
}

func getPokemonData(data *pokemonItem) *pokemonpc.Pokemon {
	return &pokemonpc.Pokemon{
		Id:          data.Id.Hex(),
		Pid:         data.Pid,
		Name:        data.Name,
		Power:       data.Power,
		Description: data.Description,
	}
}
func (*server) CreatePokemon(ctx context.Context, req *pokemonpc.CreatePokemonRequest) (*pokemonpc.CreatePokemonResponse, error) {
	fmt.Println("create pokemon")
	pokemon := req.GetPokemon()
	data := pokemonItem{
		Pid:         pokemon.GetPid(),
		Name:        pokemon.GetName(),
		Power:       pokemon.GetPower(),
		Description: pokemon.GetDescription(),
	}
	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("cannot convert to oid"),
		)
	}
	return &pokemonpc.CreatePokemonResponse{
		Pokemon: &pokemonpc.Pokemon{
			Id:          oid.Hex(),
			Pid:         pokemon.GetPid(),
			Name:        pokemon.GetName(),
			Power:       pokemon.GetPower(),
			Description: pokemon.GetDescription(),
		},
	}, nil
}
