<template>
  <v-app id="inspire">
 
    <v-app-bar app fixed clipped-left>
      <v-toolbar-title>8 Puzzle</v-toolbar-title>
    </v-app-bar>
    <v-content>
      <v-container fluid class="px-0" >
        <v-layout justify-center align-center class="px-0"  style="height:90%;">
          <v-row  style="height:100%;">
            <v-col cols="12" class="Puzzle">

              <Puzzle ref="initial" class="Puzzle__initial" :config="initial" is-initial input/>
              <span class="Puzzle__arrow">===>></span>
              <Puzzle ref="final" class="Puzzle__final" :config="final" input/>
              
            </v-col>
            <v-col cols="12" class="Puzzle">

              <v-btn :disabled="isLoading" :loading="isLoading" style="display:block;"  @click="takeMoves">
                Calculate steps
              </v-btn> 
              <v-btn  :disabled="isLoading" :loading="isLoading" style="margin-left:30px;display:block;" @click="cheat">
               Cheat
              </v-btn> 
              <v-btn :disabled="isLoading" :loading="isLoading"  style="margin-left:30px;display:block;" @click="takeStats">
               Stats
              </v-btn> 

              <v-select
              v-model="algoritm"
              :items="items"
              label="Select algoritm"
              style="margin-left:30px;display:block;max-width:200px;"
             ></v-select>
            </v-col>
         
             <v-col cols="12" v-if="movesLength" class="Puzzle">
               Steps
             </v-col>
            <v-col cols="12" style="height:100%;" class="Puzzle">
                 <v-row  style="height:100%;">
                  <v-col v-for="(matrix,index1) in list"  :key="index1" class="Puzzle" cols="12">
                    <template v-for="(move,index) in matrix" keep-alive>
                      <Puzzle  v-if="move" :config="move.item" :key="(index1 * matrix.length + index) + 'puzzle'" />
                      <span v-if=" move && move.index  !== movesLength - 1" :key="(index1 * matrix.length + index)" class="Puzzle__arrow">===>></span>
                    </template>
                  </v-col>
                 </v-row>
            </v-col>
            <v-col v-if="Math.floor(movesLength / pageSize)" cols="12" style="text-align:center;">
              <v-pagination
              v-model="page"
              :length="Math.ceil(movesLength / pageSize)"
              total-visible="20"
              :disabled="isLoadingPage"
            ></v-pagination>
            </v-col>

          </v-row>

          <v-dialog
            v-model="dialog"
            width="700"
            style="padding:20px;"
            disable-pagination
          >
           
      
            <v-card>
               <v-data-table
                :headers="headers"
                :items="stats"
                :disable-pagination="true"
                class="elevation-1"
              ></v-data-table>
                <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn
                  color="primary"
                  text
                  @click="dialog = false"
                >
                 Close
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
          
        </v-layout>
      </v-container>
    </v-content>
    <v-footer app fixed>
      <span style="margin-left:1em">&copy; Baron Tiberiu</span>
    </v-footer>
  </v-app>
</template>

<script>
import Puzzle from "./components/Puzzle.vue"

export default {
    
    components: {
      Puzzle,

    },
    props: {
      source: String
    },
    computed:{
      
    },
    watch:{
      page(newValue,oldValue){
        if(newValue != oldValue){
          this.takePage()
        }
      }
    },
    data: () => ({
      isLoading:false,
      isLoadingPage:false,
  

      page:1,
      pageSize:15,
      movesLength:0,
      dialog:false,


      algoritm:'astar',
 
      list:[],

      stats:[],
      headers: [
        {
          text: 'Algoritm',
          align: 'start',
          value: 'algoritm',
        },
        { text: 'Timp de lucru', value: 'running_time' },
        { text: 'Numarul de stari vizitate', value: 'visited_configs' },
        { text: 'Adancime stare finala', value: 'final_depth' },
        { text: 'Consume memorie', value: 'memory_usage' },
      ],

      final:{
        Move:0,
        Matrix:[
          [0,0,0],
          [0,0,0],
          [0,0,0]
        ]
      },
      initial:{
          Move:0,
          Matrix:[
            [0,0,0],
            [0,0,0],
            [0,0,0]
          ]
      },


      items:[
        {
          text:'Breadth First Search',
          value:'bfs'
        },
        {
          text:'Depth First Search',
          value:'dfs'
        },
        {
          text:'AStar',
          value:'astar'
        }
      ]
    }),

    mounted(){

      var _vm = this
      window.cheat = function(){
          _vm.cheat()
      }
      this.cheat()
    },
    methods:{

      cheat(){
        this.initial = {
          Move:0,
          Matrix:[
            [1,2,3],
            [5,6,0],
            [7,8,4]
          ]
        }

        this.final = {
          Move:0,
          Matrix:[
            [1,2,3],
            [5,8,6],
            [0,7,4]
          ]
        }
      },  
      getMatrixes(){
        const initial = this.$refs.initial.matrixToString()
        const final = this.$refs.final.matrixToString()
        return  initial + "\n" + final
      },

      validate(){
        const initialValid = this.$refs.initial.validate()
        const finalValid = this.$refs.final.validate()
        return  initialValid && finalValid 
      },

      async takeMoves(){
        if(this.isLoading || !this.validate()) return
        this.isLoading = true
        try {
          const length = await this.$backend.TakeMoves(this.algoritm,this.getMatrixes())

          this.page = 1
          this.movesLength  = length
          // this.infiniteHandler()
          this.takePage()

        }catch(e){
          this.length = 0
        }
        this.isLoading = false
      },
      async takePage(){
      
        if(this.isLoadingPage) return
          this.isLoadingPage = true
        try{

          this.list = await this.$backend.Page(this.page,this.pageSize)

        }catch(e){
          this.list = []
        }

        this.isLoadingPage = false
      },
      async takeStats(){
        if(this.isLoading || !this.validate()) return
        this.isLoading = true
        try {
          const stats = await this.$backend.Stats(this.getMatrixes())

          this.stats = stats

          this.dialog = true
        }catch(e){
          this.stats = []
        }
        this.isLoading = false
      }
    },
    
}
</script>

<style lang="scss" scoped>
  .Puzzle{
    text-align: center;
    display: flex;
    align-items: center;
    justify-content: center;

    &__initial {
      display: inline-block
    }

    &__final{
      display: inline-block;
    
    }

    &__arrow{
      display: block;
      margin-left:25px;   
      margin-right:25px;
    }
  }

.scroller {
  height: 100%;
  width: 100%;
}

.item {
  height: 32%;
  padding: 0 12px;
  display: flex;
  align-items: center;
}
</style>