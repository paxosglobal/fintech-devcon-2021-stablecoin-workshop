// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// USDKABI is the input ABI used to generate the binding from.
const USDKABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// USDKBin is the compiled bytecode used for deploying new contracts.
var USDKBin = "0x60806040523480156200001157600080fd5b506040518060400160405280600881526020017f557364546f6b656e0000000000000000000000000000000000000000000000008152506040518060400160405280600481526020017f5553444b00000000000000000000000000000000000000000000000000000000815250816003908051906020019062000096929190620001a6565b508060049080519060200190620000af929190620001a6565b505050620000d2620000c6620000d860201b60201c565b620000e060201b60201c565b620002bb565b600033905090565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b828054620001b49062000256565b90600052602060002090601f016020900481019282620001d8576000855562000224565b82601f10620001f357805160ff191683800117855562000224565b8280016001018555821562000224579182015b828111156200022357825182559160200191906001019062000206565b5b50905062000233919062000237565b5090565b5b808211156200025257600081600090555060010162000238565b5090565b600060028204905060018216806200026f57607f821691505b602082108114156200028657620002856200028c565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b611e2480620002cb6000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c8063715018a611610097578063a457c2d711610066578063a457c2d71461029d578063a9059cbb146102cd578063dd62ed3e146102fd578063f2fde38b1461032d57610100565b8063715018a61461023b5780638da5cb5b1461024557806395d89b41146102635780639dc29fac1461028157610100565b8063313ce567116100d3578063313ce567146101a157806339509351146101bf57806340c10f19146101ef57806370a082311461020b57610100565b806306fdde0314610105578063095ea7b31461012357806318160ddd1461015357806323b872dd14610171575b600080fd5b61010d610349565b60405161011a91906116ac565b60405180910390f35b61013d6004803603810190610138919061141d565b6103db565b60405161014a9190611691565b60405180910390f35b61015b6103f9565b604051610168919061184e565b60405180910390f35b61018b600480360381019061018691906113ca565b610403565b6040516101989190611691565b60405180910390f35b6101a96104fb565b6040516101b69190611869565b60405180910390f35b6101d960048036038101906101d4919061141d565b610504565b6040516101e69190611691565b60405180910390f35b6102096004803603810190610204919061141d565b6105b0565b005b6102256004803603810190610220919061135d565b61063a565b604051610232919061184e565b60405180910390f35b610243610682565b005b61024d61070a565b60405161025a9190611676565b60405180910390f35b61026b610734565b60405161027891906116ac565b60405180910390f35b61029b6004803603810190610296919061141d565b6107c6565b005b6102b760048036038101906102b2919061141d565b610850565b6040516102c49190611691565b60405180910390f35b6102e760048036038101906102e2919061141d565b61093b565b6040516102f49190611691565b60405180910390f35b6103176004803603810190610312919061138a565b610959565b604051610324919061184e565b60405180910390f35b6103476004803603810190610342919061135d565b6109e0565b005b606060038054610358906119b2565b80601f0160208091040260200160405190810160405280929190818152602001828054610384906119b2565b80156103d15780601f106103a6576101008083540402835291602001916103d1565b820191906000526020600020905b8154815290600101906020018083116103b457829003601f168201915b5050505050905090565b60006103ef6103e8610ad8565b8484610ae0565b6001905092915050565b6000600254905090565b6000610410848484610cab565b6000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600061045b610ad8565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050828110156104db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d29061176e565b60405180910390fd5b6104ef856104e7610ad8565b858403610ae0565b60019150509392505050565b60006012905090565b60006105a6610511610ad8565b84846001600061051f610ad8565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546105a191906118a0565b610ae0565b6001905092915050565b6105b8610ad8565b73ffffffffffffffffffffffffffffffffffffffff166105d661070a565b73ffffffffffffffffffffffffffffffffffffffff161461062c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106239061178e565b60405180910390fd5b6106368282610f2c565b5050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b61068a610ad8565b73ffffffffffffffffffffffffffffffffffffffff166106a861070a565b73ffffffffffffffffffffffffffffffffffffffff16146106fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106f59061178e565b60405180910390fd5b610708600061108c565b565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b606060048054610743906119b2565b80601f016020809104026020016040519081016040528092919081815260200182805461076f906119b2565b80156107bc5780601f10610791576101008083540402835291602001916107bc565b820191906000526020600020905b81548152906001019060200180831161079f57829003601f168201915b5050505050905090565b6107ce610ad8565b73ffffffffffffffffffffffffffffffffffffffff166107ec61070a565b73ffffffffffffffffffffffffffffffffffffffff1614610842576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108399061178e565b60405180910390fd5b61084c8282611152565b5050565b6000806001600061085f610ad8565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508281101561091c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109139061180e565b60405180910390fd5b610930610927610ad8565b85858403610ae0565b600191505092915050565b600061094f610948610ad8565b8484610cab565b6001905092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6109e8610ad8565b73ffffffffffffffffffffffffffffffffffffffff16610a0661070a565b73ffffffffffffffffffffffffffffffffffffffff1614610a5c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a539061178e565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610acc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac39061170e565b60405180910390fd5b610ad58161108c565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610b50576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b47906117ee565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610bc0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bb79061172e565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92583604051610c9e919061184e565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610d1b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d12906117ce565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610d8b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d82906116ce565b60405180910390fd5b610d96838383611329565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610e1c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e139061174e565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610eaf91906118a0565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610f13919061184e565b60405180910390a3610f2684848461132e565b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610f9c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f939061182e565b60405180910390fd5b610fa860008383611329565b8060026000828254610fba91906118a0565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461100f91906118a0565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051611074919061184e565b60405180910390a36110886000838361132e565b5050565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156111c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111b9906117ae565b60405180910390fd5b6111ce82600083611329565b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015611254576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161124b906116ee565b60405180910390fd5b8181036000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008282546112ab91906118f6565b92505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051611310919061184e565b60405180910390a36113248360008461132e565b505050565b505050565b505050565b60008135905061134281611dc0565b92915050565b60008135905061135781611dd7565b92915050565b60006020828403121561137357611372611a42565b5b600061138184828501611333565b91505092915050565b600080604083850312156113a1576113a0611a42565b5b60006113af85828601611333565b92505060206113c085828601611333565b9150509250929050565b6000806000606084860312156113e3576113e2611a42565b5b60006113f186828701611333565b935050602061140286828701611333565b925050604061141386828701611348565b9150509250925092565b6000806040838503121561143457611433611a42565b5b600061144285828601611333565b925050602061145385828601611348565b9150509250929050565b6114668161192a565b82525050565b6114758161193c565b82525050565b600061148682611884565b611490818561188f565b93506114a081856020860161197f565b6114a981611a47565b840191505092915050565b60006114c160238361188f565b91506114cc82611a58565b604082019050919050565b60006114e460228361188f565b91506114ef82611aa7565b604082019050919050565b600061150760268361188f565b915061151282611af6565b604082019050919050565b600061152a60228361188f565b915061153582611b45565b604082019050919050565b600061154d60268361188f565b915061155882611b94565b604082019050919050565b600061157060288361188f565b915061157b82611be3565b604082019050919050565b600061159360208361188f565b915061159e82611c32565b602082019050919050565b60006115b660218361188f565b91506115c182611c5b565b604082019050919050565b60006115d960258361188f565b91506115e482611caa565b604082019050919050565b60006115fc60248361188f565b915061160782611cf9565b604082019050919050565b600061161f60258361188f565b915061162a82611d48565b604082019050919050565b6000611642601f8361188f565b915061164d82611d97565b602082019050919050565b61166181611968565b82525050565b61167081611972565b82525050565b600060208201905061168b600083018461145d565b92915050565b60006020820190506116a6600083018461146c565b92915050565b600060208201905081810360008301526116c6818461147b565b905092915050565b600060208201905081810360008301526116e7816114b4565b9050919050565b60006020820190508181036000830152611707816114d7565b9050919050565b60006020820190508181036000830152611727816114fa565b9050919050565b600060208201905081810360008301526117478161151d565b9050919050565b6000602082019050818103600083015261176781611540565b9050919050565b6000602082019050818103600083015261178781611563565b9050919050565b600060208201905081810360008301526117a781611586565b9050919050565b600060208201905081810360008301526117c7816115a9565b9050919050565b600060208201905081810360008301526117e7816115cc565b9050919050565b60006020820190508181036000830152611807816115ef565b9050919050565b6000602082019050818103600083015261182781611612565b9050919050565b6000602082019050818103600083015261184781611635565b9050919050565b60006020820190506118636000830184611658565b92915050565b600060208201905061187e6000830184611667565b92915050565b600081519050919050565b600082825260208201905092915050565b60006118ab82611968565b91506118b683611968565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156118eb576118ea6119e4565b5b828201905092915050565b600061190182611968565b915061190c83611968565b92508282101561191f5761191e6119e4565b5b828203905092915050565b600061193582611948565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b8381101561199d578082015181840152602081019050611982565b838111156119ac576000848401525b50505050565b600060028204905060018216806119ca57607f821691505b602082108114156119de576119dd611a13565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600080fd5b6000601f19601f8301169050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60008201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206160008201527f6c6c6f77616e6365000000000000000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b611dc98161192a565b8114611dd457600080fd5b50565b611de081611968565b8114611deb57600080fd5b5056fea26469706673582212200a93fbaed9ef78b129c398b3bf614116f0f300941a2caeed8e6832c34f65599064736f6c63430008060033"

// DeployUSDK deploys a new Ethereum contract, binding an instance of USDK to it.
func DeployUSDK(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *USDK, error) {
	parsed, err := abi.JSON(strings.NewReader(USDKABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(USDKBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &USDK{USDKCaller: USDKCaller{contract: contract}, USDKTransactor: USDKTransactor{contract: contract}, USDKFilterer: USDKFilterer{contract: contract}}, nil
}

// USDK is an auto generated Go binding around an Ethereum contract.
type USDK struct {
	USDKCaller     // Read-only binding to the contract
	USDKTransactor // Write-only binding to the contract
	USDKFilterer   // Log filterer for contract events
}

// USDKCaller is an auto generated read-only Go binding around an Ethereum contract.
type USDKCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDKTransactor is an auto generated write-only Go binding around an Ethereum contract.
type USDKTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDKFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type USDKFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDKSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type USDKSession struct {
	Contract     *USDK             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// USDKCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type USDKCallerSession struct {
	Contract *USDKCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// USDKTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type USDKTransactorSession struct {
	Contract     *USDKTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// USDKRaw is an auto generated low-level Go binding around an Ethereum contract.
type USDKRaw struct {
	Contract *USDK // Generic contract binding to access the raw methods on
}

// USDKCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type USDKCallerRaw struct {
	Contract *USDKCaller // Generic read-only contract binding to access the raw methods on
}

// USDKTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type USDKTransactorRaw struct {
	Contract *USDKTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUSDK creates a new instance of USDK, bound to a specific deployed contract.
func NewUSDK(address common.Address, backend bind.ContractBackend) (*USDK, error) {
	contract, err := bindUSDK(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &USDK{USDKCaller: USDKCaller{contract: contract}, USDKTransactor: USDKTransactor{contract: contract}, USDKFilterer: USDKFilterer{contract: contract}}, nil
}

// NewUSDKCaller creates a new read-only instance of USDK, bound to a specific deployed contract.
func NewUSDKCaller(address common.Address, caller bind.ContractCaller) (*USDKCaller, error) {
	contract, err := bindUSDK(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &USDKCaller{contract: contract}, nil
}

// NewUSDKTransactor creates a new write-only instance of USDK, bound to a specific deployed contract.
func NewUSDKTransactor(address common.Address, transactor bind.ContractTransactor) (*USDKTransactor, error) {
	contract, err := bindUSDK(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &USDKTransactor{contract: contract}, nil
}

// NewUSDKFilterer creates a new log filterer instance of USDK, bound to a specific deployed contract.
func NewUSDKFilterer(address common.Address, filterer bind.ContractFilterer) (*USDKFilterer, error) {
	contract, err := bindUSDK(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &USDKFilterer{contract: contract}, nil
}

// bindUSDK binds a generic wrapper to an already deployed contract.
func bindUSDK(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(USDKABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USDK *USDKRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDK.Contract.USDKCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USDK *USDKRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDK.Contract.USDKTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USDK *USDKRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDK.Contract.USDKTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USDK *USDKCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDK.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USDK *USDKTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDK.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USDK *USDKTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDK.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_USDK *USDKCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _USDK.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_USDK *USDKSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _USDK.Contract.Allowance(&_USDK.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_USDK *USDKCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _USDK.Contract.Allowance(&_USDK.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_USDK *USDKCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _USDK.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_USDK *USDKSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _USDK.Contract.BalanceOf(&_USDK.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_USDK *USDKCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _USDK.Contract.BalanceOf(&_USDK.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_USDK *USDKCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _USDK.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_USDK *USDKSession) Decimals() (uint8, error) {
	return _USDK.Contract.Decimals(&_USDK.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_USDK *USDKCallerSession) Decimals() (uint8, error) {
	return _USDK.Contract.Decimals(&_USDK.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_USDK *USDKCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _USDK.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_USDK *USDKSession) Name() (string, error) {
	return _USDK.Contract.Name(&_USDK.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_USDK *USDKCallerSession) Name() (string, error) {
	return _USDK.Contract.Name(&_USDK.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_USDK *USDKCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDK.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_USDK *USDKSession) Owner() (common.Address, error) {
	return _USDK.Contract.Owner(&_USDK.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_USDK *USDKCallerSession) Owner() (common.Address, error) {
	return _USDK.Contract.Owner(&_USDK.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_USDK *USDKCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _USDK.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_USDK *USDKSession) Symbol() (string, error) {
	return _USDK.Contract.Symbol(&_USDK.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_USDK *USDKCallerSession) Symbol() (string, error) {
	return _USDK.Contract.Symbol(&_USDK.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_USDK *USDKCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USDK.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_USDK *USDKSession) TotalSupply() (*big.Int, error) {
	return _USDK.Contract.TotalSupply(&_USDK.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_USDK *USDKCallerSession) TotalSupply() (*big.Int, error) {
	return _USDK.Contract.TotalSupply(&_USDK.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_USDK *USDKTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDK.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_USDK *USDKSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDK.Contract.Approve(&_USDK.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_USDK *USDKTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDK.Contract.Approve(&_USDK.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_USDK *USDKTransactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDK.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_USDK *USDKSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDK.Contract.Burn(&_USDK.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_USDK *USDKTransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDK.Contract.Burn(&_USDK.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_USDK *USDKTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _USDK.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_USDK *USDKSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _USDK.Contract.DecreaseAllowance(&_USDK.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_USDK *USDKTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _USDK.Contract.DecreaseAllowance(&_USDK.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_USDK *USDKTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _USDK.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_USDK *USDKSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _USDK.Contract.IncreaseAllowance(&_USDK.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_USDK *USDKTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _USDK.Contract.IncreaseAllowance(&_USDK.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_USDK *USDKTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDK.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_USDK *USDKSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDK.Contract.Mint(&_USDK.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_USDK *USDKTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDK.Contract.Mint(&_USDK.TransactOpts, account, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_USDK *USDKTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDK.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_USDK *USDKSession) RenounceOwnership() (*types.Transaction, error) {
	return _USDK.Contract.RenounceOwnership(&_USDK.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_USDK *USDKTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _USDK.Contract.RenounceOwnership(&_USDK.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_USDK *USDKTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDK.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_USDK *USDKSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDK.Contract.Transfer(&_USDK.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_USDK *USDKTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDK.Contract.Transfer(&_USDK.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_USDK *USDKTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDK.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_USDK *USDKSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDK.Contract.TransferFrom(&_USDK.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_USDK *USDKTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDK.Contract.TransferFrom(&_USDK.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_USDK *USDKTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _USDK.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_USDK *USDKSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _USDK.Contract.TransferOwnership(&_USDK.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_USDK *USDKTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _USDK.Contract.TransferOwnership(&_USDK.TransactOpts, newOwner)
}

// USDKApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the USDK contract.
type USDKApprovalIterator struct {
	Event *USDKApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDKApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDKApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDKApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDKApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDKApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDKApproval represents a Approval event raised by the USDK contract.
type USDKApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_USDK *USDKFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*USDKApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _USDK.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &USDKApprovalIterator{contract: _USDK.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_USDK *USDKFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *USDKApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _USDK.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDKApproval)
				if err := _USDK.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_USDK *USDKFilterer) ParseApproval(log types.Log) (*USDKApproval, error) {
	event := new(USDKApproval)
	if err := _USDK.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDKOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the USDK contract.
type USDKOwnershipTransferredIterator struct {
	Event *USDKOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDKOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDKOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDKOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDKOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDKOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDKOwnershipTransferred represents a OwnershipTransferred event raised by the USDK contract.
type USDKOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_USDK *USDKFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*USDKOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _USDK.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &USDKOwnershipTransferredIterator{contract: _USDK.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_USDK *USDKFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *USDKOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _USDK.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDKOwnershipTransferred)
				if err := _USDK.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_USDK *USDKFilterer) ParseOwnershipTransferred(log types.Log) (*USDKOwnershipTransferred, error) {
	event := new(USDKOwnershipTransferred)
	if err := _USDK.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDKTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the USDK contract.
type USDKTransferIterator struct {
	Event *USDKTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *USDKTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDKTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(USDKTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *USDKTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDKTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDKTransfer represents a Transfer event raised by the USDK contract.
type USDKTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_USDK *USDKFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDKTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDK.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &USDKTransferIterator{contract: _USDK.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_USDK *USDKFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *USDKTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDK.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDKTransfer)
				if err := _USDK.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_USDK *USDKFilterer) ParseTransfer(log types.Log) (*USDKTransfer, error) {
	event := new(USDKTransfer)
	if err := _USDK.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
