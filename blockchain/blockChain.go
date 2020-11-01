package blockchain

import (
	"errors"
	"github.com/bolt-master"
	"math/big"
)

//区块链结构体实例定义
//区块链包含：
//1、将新产生的区块与已有的区块链接起来，并保存
//2、可以查询某个区块的信息
//可以将所有的区块进行遍历，输出区块信息
type BlockChain struct {
	LastHash []byte //最新区块的hash
	BoltDb   *bolt.DB
}

var BUCKET_NAME = "blocks"
var LAST_KEY = "lastbolck"
var CHAINDB = "chain.db"
var CHAIN BlockChain

//查询所有的区块信息，并返回，将所有的区块放入切片中
func (bc BlockChain) QuerAllBlocks() []*Block {
	blocks := make([]*Block, 0)

	db := bc.BoltDb
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			panic("查询数据失败")

		}
		eachKey := bc.LastHash
		preHashBig := new(big.Int)
		zeroBig := big.NewInt(0)
		for {
			eachBlockBytes := bucket.Get(eachKey)
			//反序列化以后得到的每一个区块
			eachBlocks, _ := DeSerialize(eachBlockBytes)
			//将遍历到的每一个区块结构体指针放入到[]bytes容器中
			blocks = append(blocks, eachBlocks)
			preHashBig.SetBytes(eachBlocks.PrevHash)
			if preHashBig.Cmp(zeroBig) == 0 { //通多if条件语句判断区块链遍历是否已到创世区块，如果是，跳出循环
				break
			}
			eachKey = eachBlocks.PrevHash
		}
		return nil
	})
	return blocks
}

//通过区块的高度查询某个具体的区块，返回区块实例
func (bc BlockChain) QueryBlockByHeight(height int64) *Block {
	if height < 0 { //如果目标高度小于0，则说明参数不合法
		return nil
	}
	var block *Block
	db := bc.BoltDb
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			panic("查询数据失败")

		}
		hashKey := bc.LastHash
		for {
			lastBlockByte := bucket.Get(hashKey)
			eachBlock, _ := DeSerialize(lastBlockByte)
			//给定的数字超出区块链
			if eachBlock.Height < height {
				break
			}

			if eachBlock.Height == height { //高度和目标一致
				block = eachBlock
				break
			}
			//遍历的当前的区块的高度和目标高度不一样
			hashKey = eachBlock.PrevHash

		}

		return nil
	})

	return block
}

func NewBlockChain() BlockChain {
	db, err := bolt.Open(CHAINDB, 0600, nil)
	if err != nil {
		panic(err.Error())
	}
	var bl BlockChain
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(BUCKET_NAME))
			if err != nil {
				panic(err.Error())
			}
		}
		lastHash := bucket.Get([]byte(LAST_KEY))
		if len(lastHash) == 0 { //没有创世区块
			//创建创世区块
			genesis := CreateGenesisBlock()
			//创建一个存储区块的文件

			bl = BlockChain{
				LastHash: genesis.Hash,
				BoltDb:   db,
			}
			genesisBytes, _ := genesis.Serialize()
			bucket.Put(genesis.Hash, genesisBytes)
			bucket.Put([]byte(LAST_KEY), genesis.Hash)
		} else { //有创世区块
			lastHash := bucket.Get([]byte(LAST_KEY))
			lastBlockBytes := bucket.Get(lastHash)
			lastBlock, err := DeSerialize(lastBlockBytes)
			if err != nil {
				panic("读取区块链数据失败")
			}
			bl = BlockChain{
				LastHash: lastBlock.Hash,
				BoltDb:   db,
			}
		}
		return nil
	})
	//为全局赋值
	CHAIN = bl
	return bl
}

func (bc BlockChain) QueryBlockByCertId(cert_id []byte) (*Block, *Block) {
	var block *Block
	db := bc.BoltDb
	var e error
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			e = errors.New("查询区块数据遇到错误!!!")
			return e
		}
		//桶存在
		eachKey := []byte(LAST_KEY)
		eachBig := new(big.Int)
		zeroBig := big.NewInt(0)
		for {
			eachHash := bucket.Get(eachKey)
			eachBlockBytes := bucket.Get(eachHash)
			eachBlock, _ := DeSerialize(eachBlockBytes)
			//找到的情况
			if string(eachBlock.Data) == string(cert_id) {
				block = eachBlock
				break
			}
			block = append(block, eachBlock)
			eachBig.SetBytes(eachBlock.PrevHash)
			if eachBig.Cmp(zeroBig) == 0 { //通多if条件语句判断区块链遍历是否已到创世区块，如果是，跳出循环
				break
			}
			eachKey = eachBlock.PrevHash
		}

		return nil
	})
	return nil, nil
}

//调用BlockChain的方法，该方法可以将一个生成的新区块保存到chain.db文件中
func (bc BlockChain) SaveData(data []byte) (Block, error) {
	db := bc.BoltDb
	var e error
	var lastBlock *Block
	//查询chain.db中存储最新的区块
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			panic("boltdb未创建，请重试！！！")
		}
		lastHash := bucket.Get([]byte(LAST_KEY))
		lastBlockBytes := bucket.Get(lastHash)
		lastBlock, _ = DeSerialize(lastBlockBytes)

		return nil
	})
	//生成一个区块，把data存入到新生成的区块中
	newBlock := NewBlock(lastBlock.Height+1, data, lastBlock.Hash)
	//更新chain.db
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		//序列化
		newBlockByte, _ := newBlock.Serialize()
		//把区块信息保存到boltdb中
		bucket.Put(newBlock.Hash, newBlockByte)
		//更新代表最后一个区块hash值的记录
		bucket.Put([]byte(LAST_KEY), newBlock.Hash)
		bc.LastHash = newBlock.Hash

		return nil
	})
	return newBlock, e
}
