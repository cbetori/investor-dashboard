import { connect } from "react-redux";
import FundsDistroChart from '../components/Funds/FundDistroChart'

function mapStateToProps(state) {
  return {
    distributionsFund: state.distributionsFund,
  };
}

const mapState= connect(mapStateToProps)
export const DistroFundsChartContainer = mapState(FundsDistroChart)